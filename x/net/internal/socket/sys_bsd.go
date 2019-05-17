// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build darwin dragonfly freebsd openbsd

package socket

import "errors"

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errors.New("not implemented")
=======
// +build aix darwin dragonfly freebsd openbsd

package socket

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
