package palindromeSubstrings

import (
	"fmt"
	"testing"
)

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		j, k := i, i
		for j >= 0 && k < len(s) {
			if s[j] == s[k] {
				count++
			} else {
				break

			}
			j--
			k++
		}
		if i < len(s)-1 && s[i+1] == s[i] {
			j, k = i, i+1
			for j >= 0 && k < len(s) {
				if s[j] == s[k] {
					count++
				} else {
					break

				}
				j--
				k++
			}
		}
	}
	return count
}

func TestPalidromicSubstrings(t *testing.T) {
	var test = []struct {
		in   string
		want int
	}{
		{"aaa", 6},
		{"", 0},
		{"a", 1},
		{"aa", 3},
		{"abba", 6},
		{"abc", 3},
		{"abac", 5},
		{"abbcce", 8},
		{"aabac", 7},
		{"aabbac", 9},
		{"abcba", 7},
		{"ababa", 9},
	}

	for _, tt := range test {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := countSubstrings(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
