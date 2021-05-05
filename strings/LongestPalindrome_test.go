package strings

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	start := 0
	finish := 0
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		j := i
		k := i

		for j >= 0 && k < len(chars) && chars[j] == chars[k] {
			j--
			k++
		}
		if (k - j) > (finish - start) {
			start = j + 1
			finish = k
		}
		if i < len(chars)-1 && chars[i] == chars[i+1] {
			j = i
			k = i + 1

			for j >= 0 && k < len(chars) && chars[j] == chars[k] {
				j--
				k++
			}
			if (k - j) > (finish - start) {
				start = j + 1
				finish = k
			}
		}
	}
	return string(chars[start:finish])
}

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"abccccdd", "cccc"},
		{"s", "s"},
		{"", ""},
		{"abaaaa", "aaaa"},
		{"aaaaaba", "aaaaa"},
		{"ccabbad", "abba"},
		{"ccabbadaadae", "adaada"},
		{"abcc", "cc"},
		{"ccab", "cc"},
		{"ss", "ss"},
		{"sss", "sss"},
		{"ssss", "ssss"},
		{"abba", "abba"},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := longestPalindrome(tt.in)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
