// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package socket

type msghdr struct{}

func (h *msghdr) pack(vs []iovec, bs [][]byte, oob []byte, sa []byte) {}
func (h *msghdr) name() []byte                                        { return nil }
func (h *msghdr) controllen() int                                     { return 0 }
func (h *msghdr) flags() int                                          { return 0 }
