package test

import (
	"fmt"
	"testing"
)

func LongestSubstring(s string) int {
	sArr := []rune(s)
	letters := make(map[rune]int)
	longest := 0
	for i, j := 0, 0; j < len(sArr); j++ {
		if v, ok := letters[sArr[j]]; ok && v >= i && v < j {
			i = letters[sArr[j]] + 1

		}
		if longest < (j - i + 1) {
			longest = j - i + 1
		}
		letters[sArr[j]] = j
	}
	return longest
}

// Longest substring without repeat
func TestLongestSubstring(t *testing.T) {
	var test = []struct {
		in   string
		want int
	}{
		{"abcde", 5},
		{"aaaa", 1},
		{"abba", 2},
		{"aabcdee", 5},
		{"abccc", 3},
		{"aaaab", 2},
		{"", 0},
	}

	for _, tt := range test {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := LongestSubstring(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
