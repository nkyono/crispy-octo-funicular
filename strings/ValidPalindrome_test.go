package strings

import (
	"fmt"
	"testing"
)

func ValidPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i <= j; {
		if chars[i] != chars[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func ValidPalindromeII(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i <= j; {
		if chars[i] != chars[j] {
			return ValidPalindrome(string(chars[i+1:j+1])) || ValidPalindrome(string(chars[i:j]))
		}
		i++
		j--
	}
	return true
}

func TestValidPalindrome(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		{"abc", false},
		{"aa", true},
		{"aaa", true},
		{"bbabb", true},
		{"baabb", false},
		{"", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := ValidPalindrome(tt.in)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsTwo = []struct {
		in   string
		want bool
	}{
		{"abc", false},
		{"aa", true},
		{"aaa", true},
		{"bbabb", true},
		{"baabb", true},
		{"", true},
		{"acbba", true},
		{"cabba", true},
		{"acba", true},
	}

	for _, tt := range testsTwo {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := ValidPalindromeII(tt.in)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
