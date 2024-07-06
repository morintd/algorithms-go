package mergesort

import (
	"cmp"
)

func Sort[T cmp.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	first := Sort(items[:len(items)/2])
	second := Sort(items[len(items)/2:])

	return merge(first, second)
}

func merge[T cmp.Ordered](first []T, second []T) []T {
	sorted := make([]T, 0)

	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			sorted = append(sorted, first[i])
			i++
		} else {
			sorted = append(sorted, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		sorted = append(sorted, first[i])
	}

	for ; j < len(second); j++ {
		sorted = append(sorted, second[j])
	}

	return sorted
}
