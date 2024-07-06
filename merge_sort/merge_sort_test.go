package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Run("Should return sorted array", func(t *testing.T) {
		unsorted := []int{1, 5, 3, 8, 4, 10, 15, 20, 14, 10, 11, 18}
		sorted := Sort(unsorted)
		assert.Equal(t, []int{1, 3, 4, 5, 8, 10, 10, 11, 14, 15, 18, 20}, sorted)
	})
}
