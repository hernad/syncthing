// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package rc

import (
	"os"
	"strings"

	"github.com/hernad/syncthing/lib/logger"
)

var (
	l = logger.DefaultLogger.NewFacility("rc", "Remote control package")
)

func init() {
	l.SetDebug("rc", strings.Contains(os.Getenv("STTRACE"), "rc") || os.Getenv("STTRACE") == "all")
}
