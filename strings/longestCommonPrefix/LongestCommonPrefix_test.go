package longestCommonPrefix

import (
	"fmt"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := ""
	// iterate through characters of first string (a common prefix must be contained within first string)
	for i, v := range strs[0] {
		// iterate through the rest of the strings
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || rune(strs[j][i]) != v {
				return prefix
			}
		}
		prefix = prefix + string(v)
	}
	return prefix
}

func TestLongestCommonPalindrome(t *testing.T) {
	var test = []struct {
		in   []string
		want string
	}{
		{[]string{}, ""},
		{[]string{"he", "hell", "hellscape"}, "he"},
		{[]string{"hello"}, "hello"},
		{[]string{"hellohello", "hello"}, "hello"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range test {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := longestCommonPrefix(tt.in)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
