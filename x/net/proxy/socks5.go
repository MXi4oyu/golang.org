// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proxy

import (
	"context"
	"net"
<<<<<<< HEAD

	"golang.org/x/net/internal/socks"
)

=======

	"golang.org/x/net/internal/socks"
)

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
// SOCKS5 returns a Dialer that makes SOCKSv5 connections to the given
// address with an optional username and password.
// See RFC 1928 and RFC 1929.
func SOCKS5(network, address string, auth *Auth, forward Dialer) (Dialer, error) {
	d := socks.NewDialer(network, address)
	if forward != nil {
<<<<<<< HEAD
		d.ProxyDial = func(_ context.Context, network string, address string) (net.Conn, error) {
			return forward.Dial(network, address)
		}
	}
	if auth != nil {
		up := socks.UsernamePassword{
			Username: auth.User,
			Password: auth.Password,
		}
=======
		if f, ok := forward.(ContextDialer); ok {
			d.ProxyDial = func(ctx context.Context, network string, address string) (net.Conn, error) {
				return f.DialContext(ctx, network, address)
			}
		} else {
			d.ProxyDial = func(ctx context.Context, network string, address string) (net.Conn, error) {
				return dialContext(ctx, forward, network, address)
			}
		}
	}
	if auth != nil {
		up := socks.UsernamePassword{
			Username: auth.User,
			Password: auth.Password,
		}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		d.AuthMethods = []socks.AuthMethod{
			socks.AuthMethodNotRequired,
			socks.AuthMethodUsernamePassword,
		}
		d.Authenticate = up.Authenticate
	}
	return d, nil
}
