// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build go1.9
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows

package socket

import "errors"

func (c *Conn) recvMsg(m *Message, flags int) error {
	return errors.New("not implemented")
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	return errors.New("not implemented")
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows

package socket

func (c *Conn) recvMsg(m *Message, flags int) error {
	return errNotImplemented
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
