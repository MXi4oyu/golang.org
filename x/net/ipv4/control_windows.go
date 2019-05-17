// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv4

<<<<<<< HEAD
import (
	"syscall"

	"golang.org/x/net/internal/socket"
)

func setControlMessage(c *socket.Conn, opt *rawOpt, cf ControlFlags, on bool) error {
	// TODO(mikio): implement this
	return syscall.EWINDOWS
=======
import "golang.org/x/net/internal/socket"

func setControlMessage(c *socket.Conn, opt *rawOpt, cf ControlFlags, on bool) error {
	// TODO(mikio): implement this
	return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
