// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package discover

import (
	"time"

	"github.com/hernad/syncthing/lib/protocol"
	"github.com/thejerf/suture"
)

// A Finder provides lookup services of some kind.
type Finder interface {
	Lookup(deviceID protocol.DeviceID) (direct []string, relays []Relay, err error)
	Error() error
	String() string
	Cache() map[protocol.DeviceID]CacheEntry
}

type CacheEntry struct {
	Direct     []string  `json:"direct"`
	Relays     []Relay   `json:"relays"`
	when       time.Time // When did we get the result
	found      bool      // Is it a success (cacheTime applies) or a failure (negCacheTime applies)?
	validUntil time.Time // Validity time, overrides normal calculation
}

// A FinderService is a Finder that has background activity and must be run as
// a suture.Service.
type FinderService interface {
	Finder
	suture.Service
}

type FinderMux interface {
	Finder
	ChildStatus() map[string]error
}

// The RelayStatusProvider answers questions about current relay status.
type RelayStatusProvider interface {
	Relays() []string
	RelayStatus(uri string) (time.Duration, bool)
}

// The AddressLister answers questions about what addresses we are listening
// on.
type AddressLister interface {
	ExternalAddresses() []string
	AllAddresses() []string
}
