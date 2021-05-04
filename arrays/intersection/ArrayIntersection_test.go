package test

import (
	"fmt"
	"sort"
	"testing"

	"../../sorts"
)

// Should just return the elements that are in both list (the intersection)
func SimpleArrayIntersection(arr1, arr2 []int) []int {
	arr1Set := make(map[int]bool)
	ans := make(map[int]bool)
	for _, v := range arr1 {
		arr1Set[v] = true
	}
	for _, v := range arr2 {
		if _, ok := arr1Set[v]; ok {
			ans[v] = true
		}
	}
	ret := []int{}
	for k := range ans {
		ret = append(ret, k)
	}

	sort.Slice(ret, func(a, b int) bool {
		return ret[a] < ret[b]
	})
	return ret
}

/* Same as SimpleArrayIntersection, but with restrictions
 * O(n) time and O(1) space (the resulting array of intersections is not taken into consideration).
 * Lists are sorted.
 */
func RestrictedSimpleIntersection(arr1, arr2 []int) []int {
	/*
		// only need if the two input arrays need to be sorted, I assumed input arrays were already sorted
		sort.Slice(arr1, func(a, b int) bool {
			return arr1[a] < arr1[b]
		})
		sort.Slice(arr2, func(a, b int) bool {
			return arr2[a] < arr2[b]
		})
	*/
	ans := []int{}
	if len(arr1) == 0 || len(arr2) == 0 {
		return ans
	}
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			if len(ans) == 0 || (arr1[i-1] != arr2[j-1] || arr1[i-1] != arr1[i]) {
				ans = append(ans, arr1[i])
			}
			i++
			j++
		}
	}

	return ans
}

/* returns intersection of 2 arrays
 * must return the amount of times
 * (differs from simple because simple doesn't contain duplicates)
 */

/* You could pretty easily do a merge sort on the arrays O(n*logn) + O(m*logm)
 * then you could do a O(n) iteration through using a algo similar to RestrictedSimpleIntersection
 * O(n*logn) + O(m*logm) + O(n)
 */
func ArrayIntersection(arr1, arr2 []int) []int {
	ans := []int{}
	sorts.Mergesort(arr1)
	sorts.Mergesort(arr2)

	if len(arr1) == 0 || len(arr2) == 0 {
		return ans
	}
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			ans = append(ans, arr1[i])
			i++
			j++
		}
	}

	return ans
}

// simple slice comparator, used for testing
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
		{[]int{-2, -1, 0, 1, 2}, []int{-1}, []int{-1}},
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
		testname := fmt.Sprintf("Simple Array Intersection %v, %v", tt.inArr1, tt.inArr2)
		t.Run(testname, func(t *testing.T) {
			ans := SimpleArrayIntersection(tt.inArr1, tt.inArr2)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsTwo = []struct {
		inArr1 []int
		inArr2 []int
		want   []int
	}{
		{[]int{-2, -1, 0, 1, 2}, []int{1, 2, 3}, []int{1, 2}},
		{[]int{-2, -1, 0, 1, 2}, []int{-1}, []int{-1}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5}, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{3}, []int{3}},
		{[]int{1, 13, 22, 49, 89, 91}, []int{22, 49}, []int{22, 49}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}, []int{1, 1, 1, 1, 1, 2, 2, 2, 2}, []int{1, 2}},
	}

	for _, tt := range testsTwo {
		testname := fmt.Sprintf("Restricted Simple Array Intersection %v, %v", tt.inArr1, tt.inArr2)
		t.Run(testname, func(t *testing.T) {
			ans := RestrictedSimpleIntersection(tt.inArr1, tt.inArr2)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsThree = []struct {
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
		{[]int{4, 9, 5, 9}, []int{9, 9, 4, 9}, []int{4, 9, 9}},
		{[]int{5, 6}, []int{7, 1, 5, 6}, []int{5, 6}},
		{[]int{1, 2, 2, 3, 4, 5}, []int{2, 2, 4}, []int{2, 2, 4}},
	}

	for _, tt := range testsThree {
		testname := fmt.Sprintf("Array Intersection II %v, %v", tt.inArr1, tt.inArr2)
		t.Run(testname, func(t *testing.T) {
			ans := ArrayIntersection(tt.inArr1, tt.inArr2)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
