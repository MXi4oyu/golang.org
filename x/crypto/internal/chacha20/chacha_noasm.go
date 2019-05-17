// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !s390x gccgo appengine
=======
// +build !arm64,!s390x arm64,!go1.11 gccgo appengine
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

package chacha20

const (
	bufSize = 64
	haveAsm = false
)

func (*Cipher) xorKeyStreamAsm(dst, src []byte) {
	panic("not implemented")
}
