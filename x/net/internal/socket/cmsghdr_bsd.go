// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build darwin dragonfly freebsd netbsd openbsd
=======
// +build aix darwin dragonfly freebsd netbsd openbsd
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package socket

func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = uint32(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
}
