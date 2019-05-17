// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !linux

package ipv6

import (
	"golang.org/x/net/bpf"
	"golang.org/x/net/internal/socket"
)

func (so *sockOpt) setAttachFilter(c *socket.Conn, f []bpf.RawInstruction) error {
<<<<<<< HEAD
	return errOpNoSupport
=======
	return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
