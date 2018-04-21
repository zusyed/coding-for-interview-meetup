package quickselect

func Select(a []int, k int) int {
	return quickSelect(a, 0, len(a)-1, k)
}

func quickSelect(a []int, left int, right int, k int) int {
	if left >= right {
		return a[left]
	}

	var pivot int = a[(left+right)/2]
	var index int = partition(a, left, right, pivot)

	if k == index {
		return a[k]
	} else if k < index {
		return quickSelect(a, left, index-1, k)
	} else {
		return quickSelect(a, index, right, k)
	}
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
