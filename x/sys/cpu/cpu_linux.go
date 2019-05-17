// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//+build !amd64,!amd64p32,!386

package cpu

import (
<<<<<<< HEAD
	"encoding/binary"
=======
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	"io/ioutil"
)

const (
	_AT_HWCAP  = 16
	_AT_HWCAP2 = 26

	procAuxv = "/proc/self/auxv"

<<<<<<< HEAD
	uintSize uint = 32 << (^uint(0) >> 63)
=======
	uintSize = int(32 << (^uint(0) >> 63))
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
)

// For those platforms don't have a 'cpuid' equivalent we use HWCAP/HWCAP2
// These are initialized in cpu_$GOARCH.go
// and should not be changed after they are initialized.
<<<<<<< HEAD
var HWCap uint
var HWCap2 uint
=======
var hwCap uint
var hwCap2 uint
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

func init() {
	buf, err := ioutil.ReadFile(procAuxv)
	if err != nil {
<<<<<<< HEAD
		panic("read proc auxv failed: " + err.Error())
	}

	pb := int(uintSize / 8)

	for i := 0; i < len(buf)-pb*2; i += pb * 2 {
		var tag, val uint
		switch uintSize {
		case 32:
			tag = uint(binary.LittleEndian.Uint32(buf[i:]))
			val = uint(binary.LittleEndian.Uint32(buf[i+pb:]))
		case 64:
			tag = uint(binary.LittleEndian.Uint64(buf[i:]))
			val = uint(binary.LittleEndian.Uint64(buf[i+pb:]))
		}
		switch tag {
		case _AT_HWCAP:
			HWCap = val
		case _AT_HWCAP2:
			HWCap2 = val
		}
	}
	doinit()
=======
		// e.g. on android /proc/self/auxv is not accessible, so silently
		// ignore the error and leave Initialized = false
		return
	}

	bo := hostByteOrder()
	for len(buf) >= 2*(uintSize/8) {
		var tag, val uint
		switch uintSize {
		case 32:
			tag = uint(bo.Uint32(buf[0:]))
			val = uint(bo.Uint32(buf[4:]))
			buf = buf[8:]
		case 64:
			tag = uint(bo.Uint64(buf[0:]))
			val = uint(bo.Uint64(buf[8:]))
			buf = buf[16:]
		}
		switch tag {
		case _AT_HWCAP:
			hwCap = val
		case _AT_HWCAP2:
			hwCap2 = val
		}
	}
	doinit()

	Initialized = true
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
