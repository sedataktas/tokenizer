package parser

import (
	"testing"
	"tokenizer/token"
)

func TestParseToken(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue token.Token
	}{
		{"120i8", token.Int8},
		{"1_20i8", token.Int8},
		{"1_2_0_i8", token.Int8},

		{"120i16", token.Int16},
		{"12_0i16", token.Int16},
		{"32765i16", token.Int16},

		{"1", token.Int32},
		{"+1", token.Int32},
		{"-1", token.Int32},
		{"2147483643", token.Int32},
		{"10_000", token.Int32},
		{"10_000i32", token.Int32},
		{"0x123", token.Int32},
		{"0X123", token.Int32},
		{"10_i32", token.Int32},
		{"0x1_23", token.Int32},

		{"10_000i64", token.Int64},

		{"1_20_u8", token.Uint8},
		{"1_2_0_u8", token.Uint8},
		{"254_u8", token.Uint8},

		{"1_2_0_u16", token.Uint16},
		{"1_2_0_u16", token.Uint16},
		{"65534u16", token.Uint16},

		{"10_000_000u32", token.Uint32},
		{"429496_7293u32", token.Uint32},

		{"10_0u64", token.Uint64},
		{"0x123_456u64", token.Uint64},
		{"0xa123_456u64", token.Uint64},
		{"0xA123_456u64", token.Uint64},
		{"0x429496_7293u64", token.Uint64},
		{"0x429496_7293u64", token.Uint64},

		{"4.5f32", token.Float32},
		{".14", token.Float32},
		{"._14", token.Float32},
		{"0._14", token.Float32},
		{"60_1.45", token.Float32},

		{"3.14f64", token.Float64},

		{"binalyze", token.Invalid},
		{"-_10_", token.Invalid},
		{"10_", token.Invalid},
		{"10i_32", token.Invalid},
		{"_1", token.Invalid},
		{"1_u64_", token.Invalid},
		{"129i8", token.Invalid},
		{"32768i16", token.Invalid},
		{"257_u8", token.Invalid},
		{"65536u16", token.Invalid},
		{"2147483648", token.Invalid},
		{"2147483647123123", token.Invalid},
		{"4294967296u32", token.Invalid},
		{"-3.14f64", token.Invalid},
		{"-0x123_456u64", token.Invalid},
		{"0xs123_456u64", token.Invalid},
		{"0x123_", token.Invalid},
		{"0x429496_7293u32", token.Invalid},
		{"-0x429496_7293u64", token.Invalid},
		{"10_I32", token.Invalid},
		{"0x1_23_", token.Invalid},
		{"_0x1_23_", token.Invalid},
		{"0x_1_23_", token.Invalid},
		{"3.14F64", token.Invalid},
		{"-3.14f64", token.Invalid},
		{"_3.14f64", token.Invalid},
		{"_.14", token.Invalid},
		{"-1_20_u8", token.Invalid},
		{"1_2__0_i8", token.Invalid},
		{"128i8", token.Invalid},
	}
	for i, tt := range tests {
		p := New(tt.input)
		tok := p.ParseNumber()

		if tok.String() != tt.expectedValue.String() {
			t.Fatalf("tests[%d] - token wrong. input=%s expected=%s, got=%s",
				i, tt.input, tt.expectedValue.String(), tok.String())
		}
	}
}
