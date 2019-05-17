// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !darwin,!dragonfly,!freebsd,!netbsd,!openbsd,!solaris,!windows
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!netbsd,!openbsd,!solaris,!windows
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package ipv4

import (
	"net"

	"golang.org/x/net/internal/socket"
)

func (so *sockOpt) setIPMreq(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
<<<<<<< HEAD
	return errOpNoSupport
}

func (so *sockOpt) getMulticastIf(c *socket.Conn) (*net.Interface, error) {
	return nil, errOpNoSupport
}

func (so *sockOpt) setMulticastIf(c *socket.Conn, ifi *net.Interface) error {
	return errOpNoSupport
=======
	return errNotImplemented
}

func (so *sockOpt) getMulticastIf(c *socket.Conn) (*net.Interface, error) {
	return nil, errNotImplemented
}

func (so *sockOpt) setMulticastIf(c *socket.Conn, ifi *net.Interface) error {
	return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
