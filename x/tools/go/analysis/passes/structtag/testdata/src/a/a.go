// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the test for canonical struct tags.

package a

<<<<<<< HEAD
import "encoding/xml"
=======
import (
	"a/b"
	"encoding/xml"
)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

type StructTagTest struct {
	A   int "hello"            // want "`hello` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
	B   int "\tx:\"y\""        // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag key"
	C   int "x:\"y\"\tx:\"y\"" // want "not compatible with reflect.StructTag.Get"
	D   int "x:`y`"            // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	E   int "ct\brl:\"char\""  // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
	F   int `:"emptykey"`      // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag key"
	G   int `x:"noEndQuote`    // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	H   int `x:"trunc\x0"`     // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	I   int `x:"foo",y:"bar"`  // want "not compatible with reflect.StructTag.Get: key:.value. pairs not separated by spaces"
	J   int `x:"foo"y:"bar"`   // want "not compatible with reflect.StructTag.Get: key:.value. pairs not separated by spaces"
	OK0 int `x:"y" u:"v" w:""`
	OK1 int `x:"y:z" u:"v" w:""` // note multiple colons.
	OK2 int "k0:\"values contain spaces\" k1:\"literal\ttabs\" k2:\"and\\tescaped\\tabs\""
	OK3 int `under_scores:"and" CAPS:"ARE_OK"`
}

type UnexportedEncodingTagTest struct {
	x int `json:"xx"` // want "struct field x has json tag but is not exported"
	y int `xml:"yy"`  // want "struct field y has xml tag but is not exported"
	z int
	A int `json:"aa" xml:"bb"`
}

type unexp struct{}

type JSONEmbeddedField struct {
	UnexportedEncodingTagTest `is:"embedded"`
	unexp                     `is:"embedded,notexported" json:"unexp"` // OK for now, see issue 7363
}

type AnonymousJSON struct{}
type AnonymousXML struct{}

type AnonymousJSONField struct {
	DuplicateAnonJSON int `json:"a"`

	A int "hello" // want "`hello` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
}

<<<<<<< HEAD
type DuplicateJSONFields struct {
	JSON              int `json:"a"`
	DuplicateJSON     int `json:"a"` // want "struct field DuplicateJSON repeats json tag .a. also at a.go:52"
=======
// With different names to allow using as anonymous fields multiple times.

type AnonymousJSONField2 struct {
	DuplicateAnonJSON int `json:"a"`
}
type AnonymousJSONField3 struct {
	DuplicateAnonJSON int `json:"a"`
}

type DuplicateJSONFields struct {
	JSON              int `json:"a"`
	DuplicateJSON     int `json:"a"` // want "struct field DuplicateJSON repeats json tag .a. also at a.go:64"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	IgnoredJSON       int `json:"-"`
	OtherIgnoredJSON  int `json:"-"`
	OmitJSON          int `json:",omitempty"`
	OtherOmitJSON     int `json:",omitempty"`
<<<<<<< HEAD
	DuplicateOmitJSON int `json:"a,omitempty"` // want "struct field DuplicateOmitJSON repeats json tag .a. also at a.go:52"
=======
	DuplicateOmitJSON int `json:"a,omitempty"` // want "struct field DuplicateOmitJSON repeats json tag .a. also at a.go:64"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	NonJSON           int `foo:"a"`
	DuplicateNonJSON  int `foo:"a"`
	Embedded          struct {
		DuplicateJSON int `json:"a"` // OK because it's not in the same struct type
	}
<<<<<<< HEAD
	AnonymousJSON `json:"a"` // want "struct field AnonymousJSON repeats json tag .a. also at a.go:52"

	AnonymousJSONField // want "struct field DuplicateAnonJSON repeats json tag .a. also at a.go:52"

	XML              int `xml:"a"`
	DuplicateXML     int `xml:"a"` // want "struct field DuplicateXML repeats xml tag .a. also at a.go:68"
=======
	AnonymousJSON `json:"a"` // want "struct field AnonymousJSON repeats json tag .a. also at a.go:64"

	AnonymousJSONField // want "struct field DuplicateAnonJSON repeats json tag .a. also at a.go:64"

	XML              int `xml:"a"`
	DuplicateXML     int `xml:"a"` // want "struct field DuplicateXML repeats xml tag .a. also at a.go:80"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	IgnoredXML       int `xml:"-"`
	OtherIgnoredXML  int `xml:"-"`
	OmitXML          int `xml:",omitempty"`
	OtherOmitXML     int `xml:",omitempty"`
<<<<<<< HEAD
	DuplicateOmitXML int `xml:"a,omitempty"` // want "struct field DuplicateOmitXML repeats xml tag .a. also at a.go:68"
=======
	DuplicateOmitXML int `xml:"a,omitempty"` // want "struct field DuplicateOmitXML repeats xml tag .a. also at a.go:80"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	NonXML           int `foo:"a"`
	DuplicateNonXML  int `foo:"a"`
	Embedded2        struct {
		DuplicateXML int `xml:"a"` // OK because it's not in the same struct type
	}
<<<<<<< HEAD
	AnonymousXML `xml:"a"` // want "struct field AnonymousXML repeats xml tag .a. also at a.go:68"
=======
	AnonymousXML `xml:"a"` // want "struct field AnonymousXML repeats xml tag .a. also at a.go:80"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	Attribute    struct {
		XMLName     xml.Name `xml:"b"`
		NoDup       int      `xml:"b"`                // OK because XMLName above affects enclosing struct.
		Attr        int      `xml:"b,attr"`           // OK because <b b="0"><b>0</b></b> is valid.
<<<<<<< HEAD
		DupAttr     int      `xml:"b,attr"`           // want "struct field DupAttr repeats xml attribute tag .b. also at a.go:84"
		DupOmitAttr int      `xml:"b,omitempty,attr"` // want "struct field DupOmitAttr repeats xml attribute tag .b. also at a.go:84"

		AnonymousXML `xml:"b,attr"` // want "struct field AnonymousXML repeats xml attribute tag .b. also at a.go:84"
	}

	AnonymousJSONField `json:"not_anon"` // ok; fields aren't embedded in JSON
	AnonymousJSONField `json:"-"`        // ok; entire field is ignored in JSON
=======
		DupAttr     int      `xml:"b,attr"`           // want "struct field DupAttr repeats xml attribute tag .b. also at a.go:96"
		DupOmitAttr int      `xml:"b,omitempty,attr"` // want "struct field DupOmitAttr repeats xml attribute tag .b. also at a.go:96"

		AnonymousXML `xml:"b,attr"` // want "struct field AnonymousXML repeats xml attribute tag .b. also at a.go:96"
	}

	AnonymousJSONField2 `json:"not_anon"` // ok; fields aren't embedded in JSON
	AnonymousJSONField3 `json:"-"`        // ok; entire field is ignored in JSON
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}

type UnexpectedSpacetest struct {
	A int `json:"a,omitempty"`
	B int `json:"b, omitempty"` // want "suspicious space in struct tag value"
	C int `json:"c ,omitempty"`
	D int `json:"d,omitempty, string"` // want "suspicious space in struct tag value"
	E int `xml:"e local"`
	F int `xml:"f "`                 // want "suspicious space in struct tag value"
	G int `xml:" g"`                 // want "suspicious space in struct tag value"
	H int `xml:"h ,omitempty"`       // want "suspicious space in struct tag value"
	I int `xml:"i, omitempty"`       // want "suspicious space in struct tag value"
	J int `xml:"j local ,omitempty"` // want "suspicious space in struct tag value"
	K int `xml:"k local, omitempty"` // want "suspicious space in struct tag value"
	L int `xml:" l local,omitempty"` // want "suspicious space in struct tag value"
	M int `xml:"m  local,omitempty"` // want "suspicious space in struct tag value"
	N int `xml:" "`                  // want "suspicious space in struct tag value"
	O int `xml:""`
	P int `xml:","`
	Q int `foo:" doesn't care "`
}
<<<<<<< HEAD
=======

type DuplicateWithAnotherPackage struct {
	b.AnonymousJSONField

	// The "also at" position is in a different package and directory. Use
	// "b.b" instead of "b/b" to match the relative path on Windows easily.
	DuplicateJSON int `json:"a"` // want "struct field DuplicateJSON repeats json tag .a. also at b.b.go:8"
}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
