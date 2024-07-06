package quicksort

func quicksort(unsorted []int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(unsorted, low, high)
	quicksort(unsorted, low, pivotIndex-1)
	quicksort(unsorted, pivotIndex+1, high)
}

func partition(unsorted []int, low int, high int) int {
	pivot := unsorted[high]
	currentIndex := low - 1

	for i := low; i < high; i++ {
		if unsorted[i] <= pivot {
			currentIndex++

			tmp := unsorted[i]
			unsorted[i] = unsorted[currentIndex]
			unsorted[currentIndex] = tmp
		}
	}

	currentIndex++
	unsorted[high] = unsorted[currentIndex]
	unsorted[currentIndex] = pivot

	return currentIndex
}

func Sort(unsorted []int) {
	quicksort(unsorted, 0, len(unsorted)-1)
}
