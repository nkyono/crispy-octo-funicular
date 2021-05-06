package sorts

import (
	"math/rand"
	"time"
)

func QuicksortHelper(arr []int, low, high int) int {
	// selecting a random pivot minimizes worst case complexity
	rand.Seed(time.Now().Unix())
	pivotIndex := rand.Intn(high-low) + low
	temp := arr[high]
	arr[high] = arr[pivotIndex]
	arr[pivotIndex] = temp
	// swapped arr[high] and arr[random_index] so that the random pivot is at the end of the array
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
		}
	}
	temp = arr[i]
	arr[i] = arr[high]
	arr[high] = temp
	return i
}

// sorts in place so no return val needed (could add something for error checking but it's fine)
func Quicksort(arr []int, low, high int) {
	if arr == nil {
		return
	}
	if low < high {
		partition := QuicksortHelper(arr, low, high)
		Quicksort(arr, low, partition-1)
		Quicksort(arr, partition+1, high)
	}
}
