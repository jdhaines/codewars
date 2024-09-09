package camelCase

import "testing"

var test_cases = []struct {
	in  string
	out string
}{
	{"The_Stealth_Warrior", "TheStealthWarrior"},
	{"the-Stealth-Warrior", "theStealthWarrior"},
	{"", ""},
}

func TestCamelCase(t *testing.T) {
	for _, tt := range test_cases {
		result := ToCamelCase(tt.in)
		if result != tt.out {
			t.Errorf("Expect: %s, Got %s", tt.out, result)
		} else {
			t.Logf("Pass (%s)", result)
			continue
		}
	}
}
