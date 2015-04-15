// Copyright 2015 The Sconf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sconf_test

import (
	"fmt"

	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
)

func Example() {
	var cfg = struct {
		Main struct {
			Url string
		}
	}{}
	sconf.Must(&cfg).Read(ini.File("testdata/example.ini"))
	fmt.Println(cfg.Main.Url)
	// Output: http://localhost/
}
