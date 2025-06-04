package kata_test

import "testing"

func TestRGB(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r    int
		g    int
		b    int
		want string
	}{
		{"zeros", 0, 0, 0, "000000"},
		{"123", 1, 2, 3, "010203"},
		{"white", 255, 255, 255, "FFFFFF"},
		{"gray", 254, 253, 252, "FEFDFC"},
		{"invalid", -20, 275, 125, "00FF7D"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RGB(tt.r, tt.g, tt.b)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("RGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertDecimalToHex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input int
		want  string
	}{
		{"0", 0, "00"},
		{"255", 255, "FF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertDecimalToHex(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("convertDecimalToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDigitInHex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input int
		want  string
	}{
		{"0", 0, "0"},
		{"11", 11, "B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDigitInHex(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("getDigitInHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
