package test

import (
	"fmt"
	"testing"
)

func BinarySearch(arr []int, target int) int {
	high, low := len(arr)-1, 0
	mid := (high + low) / 2

	for low < high {
		// fmt.Printf("low: %d, mid: %d, high: %d, target: %d, arr[mid]: %d\n", low, mid, high, target, arr[mid])
		if target == arr[mid] {
			return mid
		}
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (high + low) / 2
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		inArray  []int
		inTarget int
		want     int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6}, 9, -1},
		{[]int{1, 2, 3, 4, 5, 6}, -1, -1},
		{[]int{4, 5, 6}, 1, -1},
		{[]int{}, 3, -1},
		/*
			{[]int{1, 2, 3, 4, 4, 5}, 4, 4}, // I don't think I necessarily want it to take in list with multiple of the same value
			{[]int{1, 2, 2, 3, 4, 4, 5}, 2, 1},
			{[]int{2, 2, 2, 2}, 2, 0},
		*/
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d", tt.inArray, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := BinarySearch(tt.inArray, tt.inTarget)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
