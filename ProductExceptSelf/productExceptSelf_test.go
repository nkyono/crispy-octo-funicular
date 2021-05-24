package productexceptself_test

import (
	"fmt"
	"testing"
)

func prodExceptSelf(nums []int) []int {
	
	return []int{}
}
func TestProductExceptSelf(t *testing.T) {
	var tests = []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := prodExceptSelf(tt.in)
			if fmt.Sprint(ans) != fmt.Sprint(tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
