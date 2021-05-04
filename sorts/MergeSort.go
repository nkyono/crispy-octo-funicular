package sorts

func mergeHelper(arr []int, left, right, mid int) {
	leftArr := []int{}
	rightArr := []int{}
	// copy the left side elements
	for i := 0; i < mid-left+1; i++ {
		leftArr = append(leftArr, arr[left+i])
	}
	// copy the right side elements
	for i := 0; i < right-mid; i++ {
		rightArr = append(rightArr, arr[mid+i+1])
	}

	i, j, k := 0, 0, left
	// replace elements in original array with sorted elements
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	// finish moving elements in left array
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	// finish moving elements in right array
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}

}

func merge(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		merge(arr, left, mid)    // left side
		merge(arr, mid+1, right) // right side
		mergeHelper(arr, left, right, mid)
	}
}

func Mergesort(arr []int) []int {
	merge(arr, 0, len(arr)-1)
	return arr
}
