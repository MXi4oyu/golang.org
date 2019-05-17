// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package ipv6

import (
	"net"

	"golang.org/x/net/bpf"
	"golang.org/x/net/internal/socket"
)

func (so *sockOpt) getMulticastInterface(c *socket.Conn) (*net.Interface, error) {
<<<<<<< HEAD
	return nil, errOpNoSupport
}

func (so *sockOpt) setMulticastInterface(c *socket.Conn, ifi *net.Interface) error {
	return errOpNoSupport
}

func (so *sockOpt) getICMPFilter(c *socket.Conn) (*ICMPFilter, error) {
	return nil, errOpNoSupport
}

func (so *sockOpt) setICMPFilter(c *socket.Conn, f *ICMPFilter) error {
	return errOpNoSupport
}

func (so *sockOpt) getMTUInfo(c *socket.Conn) (*net.Interface, int, error) {
	return nil, 0, errOpNoSupport
=======
	return nil, errNotImplemented
}

func (so *sockOpt) setMulticastInterface(c *socket.Conn, ifi *net.Interface) error {
	return errNotImplemented
}

func (so *sockOpt) getICMPFilter(c *socket.Conn) (*ICMPFilter, error) {
	return nil, errNotImplemented
}

func (so *sockOpt) setICMPFilter(c *socket.Conn, f *ICMPFilter) error {
	return errNotImplemented
}

func (so *sockOpt) getMTUInfo(c *socket.Conn) (*net.Interface, int, error) {
	return nil, 0, errNotImplemented
}

func (so *sockOpt) setGroup(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	return errNotImplemented
}

func (so *sockOpt) setSourceGroup(c *socket.Conn, ifi *net.Interface, grp, src net.IP) error {
	return errNotImplemented
}

func (so *sockOpt) setBPF(c *socket.Conn, f []bpf.RawInstruction) error {
	return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}

func (so *sockOpt) setGroup(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	return errOpNoSupport
}

func (so *sockOpt) setSourceGroup(c *socket.Conn, ifi *net.Interface, grp, src net.IP) error {
	return errOpNoSupport
}

func (so *sockOpt) setBPF(c *socket.Conn, f []bpf.RawInstruction) error {
	return errOpNoSupport
}
