// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build s390x,!go1.11 !arm,!amd64,!s390x gccgo appengine nacl

package poly1305

// Sum generates an authenticator for msg using a one-time key and puts the
// 16-byte result into out. Authenticating two different messages with the same
// key allows an attacker to forge messages at will.
func Sum(out *[TagSize]byte, msg []byte, key *[32]byte) {
<<<<<<< HEAD
	sumGeneric(out, msg, key)
=======
	h := newMAC(key)
	h.Write(msg)
	h.Sum(out)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
