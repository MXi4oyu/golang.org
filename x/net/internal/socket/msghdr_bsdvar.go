// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build darwin dragonfly freebsd netbsd
=======
// +build aix darwin dragonfly freebsd netbsd
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package socket

func (h *msghdr) setIov(vs []iovec) {
	l := len(vs)
	if l == 0 {
		return
	}
	h.Iov = &vs[0]
	h.Iovlen = int32(l)
}
