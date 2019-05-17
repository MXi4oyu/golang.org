// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package socket

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	"net"
	"runtime"
	"unsafe"
)

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3
)

func probeProtocolStack() int {
	switch runtime.GOARCH {
	case "amd64p32", "mips64p32":
		return 4
	default:
		var p uintptr
		return int(unsafe.Sizeof(p))
	}
}

func marshalInetAddr(ip net.IP, port int, zone string) []byte {
	return nil
}

func parseInetAddr(b []byte, network string) (net.Addr, error) {
<<<<<<< HEAD
	return nil, errors.New("not implemented")
}

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	return 0, errors.New("not implemented")
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return errors.New("not implemented")
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
}

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
=======
	return nil, errNotImplemented
}

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	return 0, errNotImplemented
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return errNotImplemented
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
