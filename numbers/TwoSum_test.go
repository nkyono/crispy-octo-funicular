package numbers

import (
	"fmt"
	"testing"

	"../comparators"
)

func TwoSum(nums []int, target int) []int {
	vals := make(map[int]int)
	for i, v := range nums {
		if val, ok := vals[v]; ok {
			return []int{val, i}
		} else {
			vals[target-v] = i
		}
	}
	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		inArr    []int
		inTarget int
		want     []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %d", tt.inArr, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := TwoSum(tt.inArr, tt.inTarget)
			if !comparators.CompareSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
