package quicksort

func Sort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, left int, right int) {
	if left >= right {
		return
	}

	var pivot int = a[(left+right)/2]
	var index int = partition(a, left, right, pivot)

	quickSort(a, left, index-1)
	quickSort(a, index, right)
}

func partition(a []int, left int, right int, pivot int) int {
	for left <= right {
		for a[left] < pivot {
			left++
		}

		for a[right] > pivot {
			right--
		}

		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}

	return left
}
