// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package upgrade

import (
	"os"
	"strings"

	"github.com/hernad/syncthing/lib/logger"
)

var (
	l = logger.DefaultLogger.NewFacility("upgrade", "Binary upgrades")
)

func init() {
	l.SetDebug("upgrade", strings.Contains(os.Getenv("STTRACE"), "upgrade") || os.Getenv("STTRACE") == "all")
}
