// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package icmp

import (
	"encoding/binary"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// An Echo represents an ICMP echo request or reply message body.
type Echo struct {
	ID   int    // identifier
	Seq  int    // sequence number
	Data []byte // data
}

// Len implements the Len method of MessageBody interface.
func (p *Echo) Len(proto int) int {
	if p == nil {
		return 0
	}
	return 4 + len(p.Data)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *Echo) Marshal(proto int) ([]byte, error) {
	b := make([]byte, 4+len(p.Data))
	binary.BigEndian.PutUint16(b[:2], uint16(p.ID))
	binary.BigEndian.PutUint16(b[2:4], uint16(p.Seq))
	copy(b[4:], p.Data)
	return b, nil
}

// parseEcho parses b as an ICMP echo request or reply message body.
func parseEcho(proto int, _ Type, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 4 {
		return nil, errMessageTooShort
	}
	p := &Echo{ID: int(binary.BigEndian.Uint16(b[:2])), Seq: int(binary.BigEndian.Uint16(b[2:4]))}
	if bodyLen > 4 {
		p.Data = make([]byte, bodyLen-4)
		copy(p.Data, b[4:])
	}
	return p, nil
}

// An ExtendedEchoRequest represents an ICMP extended echo request
// message body.
type ExtendedEchoRequest struct {
	ID         int         // identifier
	Seq        int         // sequence number
	Local      bool        // must be true when identifying by name or index
	Extensions []Extension // extensions
}

// Len implements the Len method of MessageBody interface.
func (p *ExtendedEchoRequest) Len(proto int) int {
	if p == nil {
		return 0
	}
	l, _ := multipartMessageBodyDataLen(proto, false, nil, p.Extensions)
<<<<<<< HEAD
	return 4 + l
=======
	return l
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *ExtendedEchoRequest) Marshal(proto int) ([]byte, error) {
<<<<<<< HEAD
=======
	var typ Type
	switch proto {
	case iana.ProtocolICMP:
		typ = ipv4.ICMPTypeExtendedEchoRequest
	case iana.ProtocolIPv6ICMP:
		typ = ipv6.ICMPTypeExtendedEchoRequest
	default:
		return nil, errInvalidProtocol
	}
	if !validExtensions(typ, p.Extensions) {
		return nil, errInvalidExtension
	}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	b, err := marshalMultipartMessageBody(proto, false, nil, p.Extensions)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	bb := make([]byte, 4)
	binary.BigEndian.PutUint16(bb[:2], uint16(p.ID))
	bb[2] = byte(p.Seq)
	if p.Local {
		bb[3] |= 0x01
	}
	bb = append(bb, b...)
	return bb, nil
=======
	binary.BigEndian.PutUint16(b[:2], uint16(p.ID))
	b[2] = byte(p.Seq)
	if p.Local {
		b[3] |= 0x01
	}
	return b, nil
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}

// parseExtendedEchoRequest parses b as an ICMP extended echo request
// message body.
func parseExtendedEchoRequest(proto int, typ Type, b []byte) (MessageBody, error) {
<<<<<<< HEAD
	if len(b) < 4+4 {
=======
	if len(b) < 4 {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		return nil, errMessageTooShort
	}
	p := &ExtendedEchoRequest{ID: int(binary.BigEndian.Uint16(b[:2])), Seq: int(b[2])}
	if b[3]&0x01 != 0 {
		p.Local = true
	}
	var err error
<<<<<<< HEAD
	_, p.Extensions, err = parseMultipartMessageBody(proto, typ, b[4:])
=======
	_, p.Extensions, err = parseMultipartMessageBody(proto, typ, b)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	if err != nil {
		return nil, err
	}
	return p, nil
}

// An ExtendedEchoReply represents an ICMP extended echo reply message
// body.
type ExtendedEchoReply struct {
	ID     int  // identifier
	Seq    int  // sequence number
	State  int  // 3-bit state working together with Message.Code
	Active bool // probed interface is active
	IPv4   bool // probed interface runs IPv4
	IPv6   bool // probed interface runs IPv6
}

// Len implements the Len method of MessageBody interface.
func (p *ExtendedEchoReply) Len(proto int) int {
	if p == nil {
		return 0
	}
	return 4
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *ExtendedEchoReply) Marshal(proto int) ([]byte, error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint16(b[:2], uint16(p.ID))
	b[2] = byte(p.Seq)
	b[3] = byte(p.State<<5) & 0xe0
	if p.Active {
		b[3] |= 0x04
	}
	if p.IPv4 {
		b[3] |= 0x02
	}
	if p.IPv6 {
		b[3] |= 0x01
	}
	return b, nil
}

// parseExtendedEchoReply parses b as an ICMP extended echo reply
// message body.
func parseExtendedEchoReply(proto int, _ Type, b []byte) (MessageBody, error) {
	if len(b) < 4 {
		return nil, errMessageTooShort
	}
	p := &ExtendedEchoReply{
		ID:    int(binary.BigEndian.Uint16(b[:2])),
		Seq:   int(b[2]),
		State: int(b[3]) >> 5,
	}
	if b[3]&0x04 != 0 {
		p.Active = true
	}
	if b[3]&0x02 != 0 {
		p.IPv4 = true
	}
	if b[3]&0x01 != 0 {
		p.IPv6 = true
	}
	return p, nil
}
