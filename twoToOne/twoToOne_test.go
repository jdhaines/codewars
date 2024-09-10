package twoToOne

import (
	"testing"
)

var test_cases = []struct {
	in  [2]string
	out string
}{
	{[2]string{"aretheyhere", "yestheyarehere"}, "aehrsty"},
	{[2]string{"loopingisfunbutdangerous", "lessdangerousthancoding"}, "abcdefghilnoprstu"},
}

func TestTwoToOneJoshSolution(t *testing.T) {
	for _, tt := range test_cases {
		result := TwoToOne(tt.in[0], tt.in[1])
		if tt.out != result {
			t.Errorf("expected %v, got %v", tt.out, result)
		} else {
			t.Logf("expected %v, got %v", tt.out, result)
			continue
		}
	}
}

func Benchmark_Josh_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test_cases {
			TwoToOne(tt.in[0], tt.in[1])
		}
	}
}

func TestTwoToOneSmartyPantsSolution(t *testing.T) {
	for _, tt := range test_cases {
		result := TwoToOne(tt.in[0], tt.in[1])
		if tt.out != result {
			t.Errorf("expected %v, got %v", tt.out, result)
		} else {
			t.Logf("expected %v, got %v", tt.out, result)
			continue
		}
	}
}

func Benchmark_Smarty_Pants_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test_cases {
			wishIWasSmart(tt.in[0], tt.in[1])
		}
	}
}
