// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"
<<<<<<< HEAD
	"log"

	"golang.org/x/tools/internal/jsonrpc2"
)

=======

	"golang.org/x/tools/internal/jsonrpc2"
	"golang.org/x/tools/internal/lsp/xlog"
)

const defaultMessageBufferSize = 20

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
func canceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(context.Background(), "$/cancelRequest", &CancelParams{ID: *req.ID})
}

<<<<<<< HEAD
func RunClient(ctx context.Context, stream jsonrpc2.Stream, client Client, opts ...interface{}) (*jsonrpc2.Conn, Server) {
	opts = append([]interface{}{clientHandler(client), jsonrpc2.Canceler(canceller)}, opts...)
	conn := jsonrpc2.NewConn(ctx, stream, opts...)
	return conn, &serverDispatcher{Conn: conn}
}

func RunServer(ctx context.Context, stream jsonrpc2.Stream, server Server, opts ...interface{}) (*jsonrpc2.Conn, Client) {
	opts = append([]interface{}{serverHandler(server), jsonrpc2.Canceler(canceller)}, opts...)
	conn := jsonrpc2.NewConn(ctx, stream, opts...)
	return conn, &clientDispatcher{Conn: conn}
}

func sendParseError(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request, err error) {
	if _, ok := err.(*jsonrpc2.Error); !ok {
		err = jsonrpc2.NewErrorf(jsonrpc2.CodeParseError, "%v", err)
	}
	unhandledError(conn.Reply(ctx, req, nil, err))
}

// unhandledError is used in places where an error may occur that cannot be handled.
// This occurs in things like rpc handlers that are a notify, where we cannot
// reply to the caller, or in a call when we are actually attempting to reply.
// In these cases, there is nothing we can do with the error except log it, so
// we do that in this function, and the presence of this function acts as a
// useful reminder of why we are effectively dropping the error and also a
// good place to hook in when debugging those kinds of errors.
func unhandledError(err error) {
	if err == nil {
		return
	}
	log.Printf("%v", err)
=======
func NewClient(stream jsonrpc2.Stream, client Client) (*jsonrpc2.Conn, Server, xlog.Logger) {
	log := xlog.New(NewLogger(client))
	conn := jsonrpc2.NewConn(stream)
	conn.Capacity = defaultMessageBufferSize
	conn.RejectIfOverloaded = true
	conn.Handler = clientHandler(log, client)
	conn.Canceler = jsonrpc2.Canceler(canceller)
	return conn, &serverDispatcher{Conn: conn}, log
}

func NewServer(stream jsonrpc2.Stream, server Server) (*jsonrpc2.Conn, Client, xlog.Logger) {
	conn := jsonrpc2.NewConn(stream)
	client := &clientDispatcher{Conn: conn}
	log := xlog.New(NewLogger(client))
	conn.Capacity = defaultMessageBufferSize
	conn.RejectIfOverloaded = true
	conn.Handler = serverHandler(log, server)
	conn.Canceler = jsonrpc2.Canceler(canceller)
	return conn, client, log
}

func sendParseError(ctx context.Context, log xlog.Logger, conn *jsonrpc2.Conn, req *jsonrpc2.Request, err error) {
	if _, ok := err.(*jsonrpc2.Error); !ok {
		err = jsonrpc2.NewErrorf(jsonrpc2.CodeParseError, "%v", err)
	}
	if err := conn.Reply(ctx, req, nil, err); err != nil {
		log.Errorf(ctx, "%v", err)
	}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
