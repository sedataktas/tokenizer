package token

import "strconv"

type Token int

const (
	Invalid Token = iota
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
)

var tokens = map[Token]string{
	Invalid: "Invalid",
	Int8:    "signed 8bit integer",
	Int16:   "signed 16bit integer",
	Int32:   "signed 32bit integer",
	Int64:   "signed 64bit integer",
	Uint8:   "unsigned 8bit integer",
	Uint16:  "unsigned 16it integer",
	Uint32:  "unsigned 32bit integer",
	Uint64:  "unsigned 64bit integer",
	Float32: "32bit float",
	Float64: "64bit float",
}

func (tok Token) String() string {
	s := ""

	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}

	return s
}
