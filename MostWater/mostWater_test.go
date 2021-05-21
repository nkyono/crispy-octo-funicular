package mostwater_test

import (
	"fmt"
	"testing"
)

func MostWater(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := -1

	for left < right {
		currArea := (right - left)
		if height[left] < height[right] {
			currArea = currArea * height[left]
			left++
		} else {
			currArea = currArea * height[right]
			right--
		}
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}

func TestMostWater(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := MostWater(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
