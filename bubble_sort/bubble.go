package bubblesort

func Sort(unsorted []int) {
	for i := 0; i < len(unsorted); i++ {
		for j := 0; j < len(unsorted)-1; j++ {
			if unsorted[j] > unsorted[j+1] {
				unsorted[j] += unsorted[j+1]
				unsorted[j+1] = unsorted[j] - unsorted[j+1]
				unsorted[j] -= unsorted[j+1]
			}
		}
	}
}
