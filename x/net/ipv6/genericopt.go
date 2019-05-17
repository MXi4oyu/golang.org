// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv6

// TrafficClass returns the traffic class field value for outgoing
// packets.
func (c *genericOpt) TrafficClass() (int, error) {
	if !c.ok() {
		return 0, errInvalidConn
	}
	so, ok := sockOpts[ssoTrafficClass]
	if !ok {
<<<<<<< HEAD:x/net/ipv6/genericopt.go
		return 0, errOpNoSupport
=======
		return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv6/genericopt.go
	}
	return so.GetInt(c.Conn)
}

// SetTrafficClass sets the traffic class field value for future
// outgoing packets.
func (c *genericOpt) SetTrafficClass(tclass int) error {
	if !c.ok() {
		return errInvalidConn
	}
	so, ok := sockOpts[ssoTrafficClass]
	if !ok {
<<<<<<< HEAD:x/net/ipv6/genericopt.go
		return errOpNoSupport
=======
		return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv6/genericopt.go
	}
	return so.SetInt(c.Conn, tclass)
}

// HopLimit returns the hop limit field value for outgoing packets.
func (c *genericOpt) HopLimit() (int, error) {
	if !c.ok() {
		return 0, errInvalidConn
	}
	so, ok := sockOpts[ssoHopLimit]
	if !ok {
<<<<<<< HEAD:x/net/ipv6/genericopt.go
		return 0, errOpNoSupport
=======
		return 0, errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv6/genericopt.go
	}
	return so.GetInt(c.Conn)
}

// SetHopLimit sets the hop limit field value for future outgoing
// packets.
func (c *genericOpt) SetHopLimit(hoplim int) error {
	if !c.ok() {
		return errInvalidConn
	}
	so, ok := sockOpts[ssoHopLimit]
	if !ok {
<<<<<<< HEAD:x/net/ipv6/genericopt.go
		return errOpNoSupport
=======
		return errNotImplemented
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a:x/net/ipv6/genericopt.go
	}
	return so.SetInt(c.Conn, hoplim)
}
