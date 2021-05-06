package sorts

import (
	"fmt"
	"testing"

	"../comparators"
)

// to run tests run: go test -v sorts/sorts_test.go sorts/MergeSort.go
// need to run all files together
// go test -v sorts/*.go
func TestMergesort(t *testing.T) {
	var tests = []struct {
		in   []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 1, 2}, []int{1, 1, 2, 2}},
		{[]int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
		{[]int{8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := Mergesort(tt.in)
			if !comparators.CompareSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	var tests = []struct {
		in   []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 1, 2}, []int{1, 1, 2, 2}},
		{[]int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
		{[]int{8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			Quicksort(tt.in, 0, len(tt.in)-1)
			if !comparators.CompareSlice(tt.in, tt.want) {
				t.Errorf("got %v, want %v", tt.in, tt.want)
			}
		})
	}
}
