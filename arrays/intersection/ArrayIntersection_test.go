package test

import (
	"fmt"
	"sort"
	"testing"
)

func SimpleArrayIntersection(arr1, arr2 []int) []int {
	arr1Set := make(map[int]bool)
	arr2Set := make(map[int]bool)
	ans := []int{}
	for _, v := range arr1 {
		arr1Set[v] = true
	}
	for _, v := range arr2 {
		arr2Set[v] = true
	}
	for k := range arr1Set {
		if _, ok := arr2Set[k]; ok {
			ans = append(ans, k)
		}
	}
	// have to sort due to the fact that maps may not return keys in same order everytime
	sort.Slice(ans, func(a, b int) bool {
		return ans[a] < ans[b]
	})
	return ans
}

func compSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

/*	[1,2,3,4,5], [2,3,4] ==> [2,3,4]
 *	a[0] == b[0], a++, b++
 *	a[0] != b[0] ??? => which should we iterate
 *	Dynamic Programming, if we keep an array
 */

func TestArrayIntersection(t *testing.T) {
	/* First test will test the simple array intersection
	 * Should just return the elements that are in both list (the intersection)
	 * No limitations on time or space
	 */
	var testsOne = []struct {
		inArr1 []int
		inArr2 []int
		want   []int
	}{
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5}, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{3}, []int{3}},
		{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{2, 7, 1, 4}, []int{1, 2, 7}},
		{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{5, 2, 7, 1, 4}, []int{1, 2, 5, 7}},
		{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{5, 2, 7, 1}, []int{1, 2, 5, 7}},
		{[]int{4, 9, 5, 9}, []int{9, 9, 4, 9}, []int{4, 9}},
		{[]int{5, 6}, []int{7, 1, 5, 6}, []int{5, 6}},
	}

	for _, tt := range testsOne {
		testname := fmt.Sprintf("Array Intersection II %v, %v", tt.inArr1, tt.inArr2)
		t.Run(testname, func(t *testing.T) {
			ans := SimpleArrayIntersection(tt.inArr1, tt.inArr2)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	/*
		var testsTwo = []struct {
			inArr1 []int
			inArr2 []int
			want   []int
		}{
			{[]int{}, []int{1, 2, 3}, []int{}},
			{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
			{[]int{}, []int{}, []int{}},
			{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3}, []int{1, 2, 3}},
			{[]int{1, 2, 3, 4, 5}, []int{4, 5}, []int{4, 5}},
			{[]int{1, 2, 3, 4, 5}, []int{3}, []int{3}},
			{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{2, 7, 1, 4}, []int{2, 7, 1}},
			{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{5, 2, 7, 1, 4}, []int{2, 7, 1}},
			{[]int{5, 6, 2, 7, 1, 3, 0, 9}, []int{5, 2, 7, 1}, []int{2, 7, 1}},
			{[]int{4, 9, 5, 9}, []int{9, 9, 4, 9}, []int{4, 9}},
			{[]int{5, 6}, []int{7, 1, 5, 6}, []int{5, 6}},
		}

		for _, tt := range testsTwo {
			testname := fmt.Sprintf("Array Intersection II %v, %v", tt.inArr1, tt.inArr2)
			t.Run(testname, func(t *testing.T) {
				ans := SimpleArrayIntersection(tt.inArr1, tt.inArr2)
				if !compSlice(ans, tt.want) {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			})
		}
	*/
}
