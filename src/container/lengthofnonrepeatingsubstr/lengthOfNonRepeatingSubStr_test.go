package main

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		{"abcabc", 3},
		{"bbbb", 1},
		{"中国人abc", 6},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; expected %d", tt.s, actual, tt.ans)
		}
	}
}

/**
	性能测试
 */
func BenchmarkSubstr(b *testing.B) {
	s := "中国人abc"
	ans := 6
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; expected %d", s, actual, ans)
		}
	}
}
