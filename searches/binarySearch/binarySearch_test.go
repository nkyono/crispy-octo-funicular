package test

import (
	"fmt"
	"testing"
)

func RangeBinarySearch(arr []int, left, right, target int) int {
	if left < 0 || right > len(arr) || len(arr) == 0 {
		return -1
	}
	if left == right && arr[left] == target {
		return left
	}
	// fmt.Printf("%d, %v, %d, %d\n", target, arr, left, right)
	high, low := right, left
	mid := (high + low) / 2

	for low <= high {
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

func DoubleBinarySearch(arr []int, target int) []int {
	ans := []int{-1, -1}
	mid := BinarySearch(arr, target)
	if mid != -1 {
		temp := -1
		left := mid
		right := mid
		for left != -1 {
			temp = left
			left = RangeBinarySearch(arr, 0, left-1, target)
		}
		ans[0] = temp
		for right != -1 {
			temp = right
			right = RangeBinarySearch(arr, right+1, len(arr)-1, target)
		}
		ans[1] = temp
	}
	return ans
}

// this doesn't actually work if there is only a singular instance of the target
func SpecialBinarySearch(arr []int, target int) []int {
	high, low := len(arr)-1, 0
	mid := (high + low) / 2

	/* Actually a much better strategy is to use binary search on each side of the found index
	 * If that search returned -1, we would know that it's the furthest most index
	 * by doing multiple binary searches we would stay in an O(log(n)) range rather than the O(n) case here
	 * worst case is that array only contains target so we would iterate through the whole array
	 * binary would continuously divide in half leading to desired O(log(n)) behavior
	 */
	for low < high {
		if target == arr[mid] {
			i := 0
			// works since mid-i is first check so guaranteed that mid-(i+1) will only happen when mid-i > 0
			for mid-i > 0 && arr[mid-(i+1)] == target {
				i++
			}
			j := 0
			for mid+j < len(arr)-1 && arr[mid+(j+1)] == target {
				j++
			}
			return []int{mid - i, mid + j}
		}
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (high + low) / 2
	}

	return []int{-1, -1}
}

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
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Regular: %v, %d", tt.inArray, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := BinarySearch(tt.inArray, tt.inTarget)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	var testsRange = []struct {
		inArray  []int
		inTarget int
		left     int
		right    int
		want     int
	}{
		{[]int{1, 2, 4, 5, 6}, 4, 0, 4, 2},
		{[]int{1, 2, 4, 5, 6, 7}, 4, 0, 3, 2},
		{[]int{1, 2, 4, 5, 6, 7}, 6, 4, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6}, 2, 2, 5, -1},
		{[]int{4, 5, 6}, 1, 0, 2, -1},
		{[]int{}, 3, 0, 0, -1},
	}

	for _, tt := range testsRange {
		testname := fmt.Sprintf("Range: %v, %d", tt.inArray, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := RangeBinarySearch(tt.inArray, tt.left, tt.right, tt.inTarget)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsSpecial = []struct {
		inArray  []int
		inTarget int
		want     []int
	}{
		{[]int{1, 2, 4, 4, 5, 6}, 4, []int{2, 3}},
		{[]int{1, 1, 1, 2, 4, 5, 6}, 1, []int{0, 2}},
		{[]int{1, 2, 4, 4, 4}, 4, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 9, []int{-1, -1}},
		{[]int{1, 2, 3, 4, 5, 6}, -1, []int{-1, -1}},
		{[]int{4, 5, 6}, 1, []int{-1, -1}},
		{[]int{}, 3, []int{-1, -1}},
		{[]int{1, 1, 1, 1}, 1, []int{0, 3}},
	}

	for _, tt := range testsSpecial {
		testname := fmt.Sprintf("Special: %v, %d", tt.inArray, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := SpecialBinarySearch(tt.inArray, tt.inTarget)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsDouble = []struct {
		inArray  []int
		inTarget int
		want     []int
	}{
		{[]int{1, 2, 4, 4, 5, 6}, 4, []int{2, 3}},
		{[]int{1, 1, 1, 2, 4, 5, 6}, 1, []int{0, 2}},
		{[]int{1, 2, 4, 4, 4}, 4, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 9, []int{-1, -1}},
		{[]int{1, 2, 3, 4, 5, 6}, -1, []int{-1, -1}},
		{[]int{4, 5, 6}, 1, []int{-1, -1}},
		{[]int{}, 3, []int{-1, -1}},
		{[]int{1, 1, 1, 1}, 1, []int{0, 3}},
	}

	for _, tt := range testsDouble {
		testname := fmt.Sprintf("Double: %v, %d", tt.inArray, tt.inTarget)
		t.Run(testname, func(t *testing.T) {
			ans := DoubleBinarySearch(tt.inArray, tt.inTarget)
			if !compSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
