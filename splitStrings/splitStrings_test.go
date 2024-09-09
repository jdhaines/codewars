package splitstrings

import (
	"reflect"
	"testing"
)

var test_cases = []struct {
	in  string
	out []string
}{
	{"abc", []string{"ab", "c_"}},
	{"dawsd", []string{"da", "ws", "d_"}},
	{"awsaws", []string{"aw", "sa", "ws"}},
	{"drr", []string{"dr", "r_"}},
	{"golangisawesome", []string{"go", "la", "ng", "is", "aw", "es", "om", "e_"}},
}

func TestSplitStrings(t *testing.T) {
	for _, tt := range test_cases {
		result := Solution(tt.in)
		if !reflect.DeepEqual(tt.out, result) {
			t.Errorf("expected %v, got %v", tt.out, result)
		} else {
			t.Logf("expected %v, got %v", tt.out, result)
			continue
		}
	}
}
