// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package scanner

import (
	"os"
	"strings"

	"github.com/hernad/syncthing/lib/logger"
)

var (
	l = logger.DefaultLogger.NewFacility("scanner", "File change detection and hashing")
)

func init() {
	l.SetDebug("scanner", strings.Contains(os.Getenv("STTRACE"), "scanner") || os.Getenv("STTRACE") == "all")
}
