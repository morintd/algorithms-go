package binarysearchlist

func Search(haystack []int, value int) int {
	for low, high := 0, len(haystack); low < high; {
		middle := (low + (high-low)/2)
		current := haystack[middle]

		if current == value {
			return middle
		} else if current > value {
			high = middle
		} else {
			low = middle + 1
		}

	}

	return -1
}
