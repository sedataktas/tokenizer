package parser

import (
	"regexp"
	"strconv"
	"strings"
	"tokenizer/token"
)

var (
	uintRgx  *regexp.Regexp
	floatRgx *regexp.Regexp
	intRgx   *regexp.Regexp
	hexRgx   *regexp.Regexp
)

type Parser struct {
	input string
}

func init() {
	var err error
	uintRgx, err = regexp.Compile(`^(\d+[_]?)*$`)
	if err != nil {
		panic(err)
	}

	intRgx, err = regexp.Compile(`^[-+]?(\d+[_]?)*$`)
	if err != nil {
		panic(err)
	}

	floatRgx, err = regexp.Compile(`^([_]?\d+[_]?)*$`)
	if err != nil {
		panic(err)
	}

	hexRgx, err = regexp.Compile(`^([0-9a-fA-F]+[_]?)*$`)
	if err != nil {
		panic(err)
	}
}

func New(input string) *Parser {
	p := &Parser{input: input}
	return p
}

func (p *Parser) ParseNumber() token.Token {
	in := p.input
	lower := strings.ToLower(in)

	if strings.Contains(in, ".") {
		return parseFloat(in)
	} else if strings.Contains(lower, "x") {
		return parseHex(lower)
	} else {
		return parseInt(in)
	}
}

func parseFloat(in string) token.Token {
	// Float parse : must contain "."
	// 4.5f32 || .14 || 60_1.45 || 3.14f64 ...
	s := strings.Split(in, ".")
	if len(s) > 2 {
		return token.Invalid
	}

	if len(s) == 2 {
		// s[0] : 4 || "" || 60_1 || 3
		if uintRgx.MatchString(s[0]) {
			// s[1] : 5f32 || 14 || 45 || 14f64
			return getFloat(s[1])
		}
	}
	return token.Invalid
}

func parseHex(lower string) token.Token {
	// Hex Parse : must contain "x" or "X"
	// 0x123 || 0X123 || 0xA12u32 || 0x123_456u64 ...
	s := strings.Split(lower, "x")
	if len(s) > 2 {
		return token.Invalid
	}

	if len(s) == 2 {
		// s[0] : must contain "0"
		if s[0] == "0" {
			// s[1] : 123 || 123 || A12u32 || 123_456u64
			return getHex(s[1])
		}
	}
	return token.Invalid
}

func parseInt(in string) token.Token {
	return getInt(in)
}

func getFloat(in string) token.Token {
	switch true {
	case strings.Contains(in, "f"):
		// 4.5f32
		s := strings.Split(in, "f")
		if len(s) > 2 {
			return token.Invalid
		}

		// s[0] : 4.5
		// s[1] : 32
		if floatRgx.MatchString(s[0]) {
			s[0] = clearUnderscores(s[0])

			if s[1] == "32" {
				if isFloatValid(s[0], 32) {
					return token.Float32
				}
				return token.Invalid
			}

			if s[1] == "64" {
				if isFloatValid(s[0], 64) {
					return token.Float64
				}
				return token.Invalid
			}
		}

		return token.Invalid
	default:
		s := strings.Split(in, "")

		lastChar := len(s) - 1
		if s[lastChar] == "_" {
			return token.Invalid
		}

		if floatRgx.MatchString(in) {
			in = clearUnderscores(in)

			if isFloatValid(in, 32) {
				return token.Float32
			}
		}
		return token.Invalid
	}
}

func getHex(in string) token.Token {
	switch true {
	case strings.Contains(in, "u"):
		// A12u32
		s := strings.Split(in, "u")
		if len(s) > 2 {
			return token.Invalid
		}

		// split[0] : A12
		// split[1] : 32
		if hexRgx.MatchString(s[0]) {
			s[0] = clearUnderscores(s[0])
			if s[1] == "32" {
				if isHexValid(s[0], 32) {
					return token.Uint32
				}
				return token.Invalid
			}

			if s[1] == "64" {
				if isHexValid(s[0], 64) {
					return token.Uint64
				}
				return token.Invalid
			}
		}
		return token.Invalid
	default:
		s := strings.Split(in, "")

		lastChar := len(s) - 1
		if s[lastChar] == "_" {
			return token.Invalid
		}
		if hexRgx.MatchString(in) {
			in = clearUnderscores(in)
			if isHexValid(in, 32) {
				return token.Int32
			}
		}
		return token.Invalid
	}
}

func getInt(in string) token.Token {
	switch true {
	case strings.Contains(in, "i"):
		// 1_20_i32
		s := strings.Split(in, "i")
		if len(s) > 2 {
			return token.Invalid
		}

		// s[0] : 1_20_
		// s[1] : 32
		if intRgx.MatchString(s[0]) {
			s[0] = clearUnderscores(s[0])

			if s[1] == "8" {
				if isIntValid(s[0], 8) {
					return token.Int8
				}
				return token.Invalid
			}

			if s[1] == "16" {
				if isIntValid(s[0], 16) {
					return token.Int16
				}
				return token.Invalid
			}

			if s[1] == "32" {
				if isIntValid(s[0], 32) {
					return token.Int32
				}
				return token.Invalid
			}

			if s[1] == "64" {
				if isIntValid(s[0], 64) {
					return token.Int64
				}
			}
		}

		return token.Invalid
	case strings.Contains(in, "u"):
		// 1_20_u8
		s := strings.Split(in, "u")
		if len(s) > 2 {
			return token.Invalid
		}

		// s[0] : 1_20_
		// s[1] : 8
		if uintRgx.MatchString(s[0]) {
			s[0] = clearUnderscores(s[0])
			if s[1] == "8" {
				if isUintValid(s[0], 8) {
					return token.Uint8
				}
			}

			if s[1] == "16" {
				if isUintValid(s[0], 16) {
					return token.Uint16
				}
			}

			if s[1] == "32" {
				if isUintValid(s[0], 32) {
					return token.Uint32
				}
			}

			if s[1] == "64" {
				if isUintValid(s[0], 64) {
					return token.Uint64
				}
			}
		}
		return token.Invalid
	default:
		// split : ["1", "_"]
		s := strings.Split(in, "")

		// cannot contain "_" end of int : "1_"
		lastChar := len(s) - 1
		if s[lastChar] == "_" {
			return token.Invalid
		}

		if intRgx.MatchString(in) {
			in = clearUnderscores(in)

			if isIntValid(in, 32) {
				return token.Int32
			}
		}
		return token.Invalid
	}
}

func isIntValid(s string, bitSize int) bool {
	_, err := strconv.ParseInt(s, 10, bitSize)
	return err == nil
}

func isUintValid(s string, bitSize int) bool {
	_, err := strconv.ParseUint(s, 10, bitSize)
	return err == nil
}

func isHexValid(s string, bitSize int) bool {
	_, err := strconv.ParseUint(s, 16, bitSize)
	return err == nil
}

func isFloatValid(s string, bitSize int) bool {
	_, err := strconv.ParseFloat(s, bitSize)
	return err == nil
}

func clearUnderscores(s string) string {
	return strings.ReplaceAll(s, "_", "")
}
