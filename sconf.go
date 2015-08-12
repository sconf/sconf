// Copyright 2015 The Sconf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sconf // import "gopkg.in/sconf/sconf.v0"

import (
	"gopkg.in/sconf/internal.v0/internal-"
	"gopkg.in/sconf/internal.v0/internal-/gcfg"
)

func Must(ptr interface{}) internal.Struct {
	_ = gcfg.ReadStringInto(ptr, "") // may panic but no parse error
	return internal.Struct{Ptr: ptr}
}
