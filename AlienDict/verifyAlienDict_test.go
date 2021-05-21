package aliendict_test

import (
	"fmt"
	"testing"
)

func orderToMap(order string) map[rune]int {
	orderMap := make(map[rune]int)
	for i, v := range order {
		orderMap[v] = i
	}
	return orderMap
}

func isWordsorted(word string, orderMap map[rune]int) bool {
	if len(word) == 1 {
		return true
	}
	for i := 1; i < len(word); i++ {
		if orderMap[rune(word[i])] < orderMap[rune(word[i-1])] {
			return false
		}
	}
	return true
}

func isAlienSorted(words []string, order string) bool {
	if len(words) == 0 {
		return true
	}
	orderMap := orderToMap(order)
	for _, v := range words {
		if !isWordsorted(v, orderMap) {
			return false
		}
	}
	return true
}

// totally misread the problem
func TestAlienDict(t *testing.T) {
	var tests = []struct {
		inWords []string
		inOrder string
		want    bool
	}{
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.inWords)
		t.Run(testname, func(t *testing.T) {
			ans := isAlienSorted(tt.inWords, tt.inOrder)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
