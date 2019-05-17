// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build freebsd netbsd openbsd
=======
// +build aix freebsd netbsd openbsd
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package socket

import (
	"runtime"
	"unsafe"
)

func probeProtocolStack() int {
<<<<<<< HEAD
	if runtime.GOOS == "openbsd" && runtime.GOARCH == "arm" {
		return 8
	}
=======
	if (runtime.GOOS == "netbsd" || runtime.GOOS == "openbsd") && runtime.GOARCH == "arm" {
		return 8
	}
	if runtime.GOOS == "aix" {
		return 1
	}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	var p uintptr
	return int(unsafe.Sizeof(p))
}
