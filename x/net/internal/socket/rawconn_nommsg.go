// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build go1.9
=======
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
// +build !linux

package socket

<<<<<<< HEAD
import "errors"

func (c *Conn) recvMsgs(ms []Message, flags int) (int, error) {
	return 0, errors.New("not implemented")
}

func (c *Conn) sendMsgs(ms []Message, flags int) (int, error) {
	return 0, errors.New("not implemented")
=======
func (c *Conn) recvMsgs(ms []Message, flags int) (int, error) {
	return 0, errNotImplemented
}

func (c *Conn) sendMsgs(ms []Message, flags int) (int, error) {
	return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
