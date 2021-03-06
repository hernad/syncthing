// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"container/heap"
	"fmt"

	"github.com/hernad/syncthing/lib/db"
	"github.com/hernad/syncthing/lib/protocol"
	"github.com/syndtr/goleveldb/leveldb"
)

// An IntHeap is a min-heap of ints.
type SizedElement struct {
	key  string
	size int
}

type ElementHeap []SizedElement

func (h ElementHeap) Len() int           { return len(h) }
func (h ElementHeap) Less(i, j int) bool { return h[i].size > h[j].size }
func (h ElementHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ElementHeap) Push(x interface{}) {
	*h = append(*h, x.(SizedElement))
}

func (h *ElementHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dumpsize(ldb *leveldb.DB) {
	h := &ElementHeap{}
	heap.Init(h)

	it := ldb.NewIterator(nil, nil)
	var dev protocol.DeviceID
	var ele SizedElement
	for it.Next() {
		key := it.Key()
		switch key[0] {
		case db.KeyTypeDevice:
			folder := nulString(key[1 : 1+64])
			devBytes := key[1+64 : 1+64+32]
			name := nulString(key[1+64+32:])
			copy(dev[:], devBytes)
			ele.key = fmt.Sprintf("DEVICE:%s:%s:%s", dev, folder, name)

		case db.KeyTypeGlobal:
			folder := nulString(key[1 : 1+64])
			name := nulString(key[1+64:])
			ele.key = fmt.Sprintf("GLOBAL:%s:%s", folder, name)

		case db.KeyTypeBlock:
			folder := nulString(key[1 : 1+64])
			hash := key[1+64 : 1+64+32]
			name := nulString(key[1+64+32:])
			ele.key = fmt.Sprintf("BLOCK:%s:%x:%s", folder, hash, name)

		case db.KeyTypeDeviceStatistic:
			ele.key = fmt.Sprintf("DEVICESTATS:%s", key[1:])

		case db.KeyTypeFolderStatistic:
			ele.key = fmt.Sprintf("FOLDERSTATS:%s", key[1:])

		case db.KeyTypeVirtualMtime:
			ele.key = fmt.Sprintf("MTIME:%s", key[1:])

		default:
			ele.key = fmt.Sprintf("UNKNOWN:%x", key)
		}
		ele.size = len(it.Value())
		heap.Push(h, ele)
	}

	for h.Len() > 0 {
		ele = heap.Pop(h).(SizedElement)
		fmt.Println(ele.key, ele.size)
	}
}
