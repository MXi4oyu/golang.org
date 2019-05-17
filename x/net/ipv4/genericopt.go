// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv4

// TOS returns the type-of-service field value for outgoing packets.
func (c *genericOpt) TOS() (int, error) {
	if !c.ok() {
		return 0, errInvalidConn
	}
	so, ok := sockOpts[ssoTOS]
	if !ok {
<<<<<<< HEAD:x/net/ipv4/genericopt.go
		return 0, errOpNoSupport
=======
		return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv4/genericopt.go
	}
	return so.GetInt(c.Conn)
}

// SetTOS sets the type-of-service field value for future outgoing
// packets.
func (c *genericOpt) SetTOS(tos int) error {
	if !c.ok() {
		return errInvalidConn
	}
	so, ok := sockOpts[ssoTOS]
	if !ok {
<<<<<<< HEAD:x/net/ipv4/genericopt.go
		return errOpNoSupport
=======
		return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv4/genericopt.go
	}
	return so.SetInt(c.Conn, tos)
}

// TTL returns the time-to-live field value for outgoing packets.
func (c *genericOpt) TTL() (int, error) {
	if !c.ok() {
		return 0, errInvalidConn
	}
	so, ok := sockOpts[ssoTTL]
	if !ok {
<<<<<<< HEAD:x/net/ipv4/genericopt.go
		return 0, errOpNoSupport
=======
		return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv4/genericopt.go
	}
	return so.GetInt(c.Conn)
}

// SetTTL sets the time-to-live field value for future outgoing
// packets.
func (c *genericOpt) SetTTL(ttl int) error {
	if !c.ok() {
		return errInvalidConn
	}
	so, ok := sockOpts[ssoTTL]
	if !ok {
<<<<<<< HEAD:x/net/ipv4/genericopt.go
		return errOpNoSupport
=======
		return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv4/genericopt.go
	}
	return so.SetInt(c.Conn, ttl)
}
