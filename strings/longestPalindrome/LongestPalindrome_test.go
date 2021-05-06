package longestPalindrome

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) int {
	// if string is empty or of length 1, we know the longest palindrome is 0 or 1 respectively
	if len(s) <= 1 {
		return len(s)
	}
	// we can count # of times a character appears, we know that if a character appears an even # of times that it can be added to palindrome
	// a palindrome can be odd surrounding a center or even with a two of the same characters in center
	// yyyyyxyyyyy or yyxxyy etc.
	// there can only be one character that appears once, if string is abcc, then only one of a or b can appear in palindrome
	// if there is an odd number of a character that appears n > 1 we can add n - 1 to the palindrome
	charCount := make(map[rune]int)
	for _, c := range s {
		if v, ok := charCount[c]; ok {
			charCount[c] = v + 1
		} else {
			charCount[c] = 1
		}
	}
	// flip the map to make it count -> # times count appears
	// to be fair, I didn't need to do this part (could have just used same logic or charCount), but I wanted to
	nums := make(map[int]int)
	for _, count := range charCount {
		if v, ok := nums[count]; ok {
			nums[count] = v + 1
		} else {
			nums[count] = 1
		}
	}
	palindromeLength := 0
	if _, ok := nums[1]; ok {
		palindromeLength++
	}
	for k, v := range nums {
		if k%2 == 0 {
			palindromeLength = palindromeLength + (k * v)
		} else {
			if palindromeLength%2 == 0 {
				palindromeLength = palindromeLength + k
				v = v - 1
			}
			palindromeLength = palindromeLength + ((k - 1) * v)
		}
	}

	return palindromeLength
}

func TestLongestPalindrome(t *testing.T) {
	var test = []struct {
		in   string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
		{"bbb", 3},
		{"", 0},
		{"ab", 1},
		{"abb", 3},
		{"bba", 3},
		{"bbaaa", 5},
		{"aaabaaaa", 7},
	}

	for _, tt := range test {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := longestPalindrome(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
