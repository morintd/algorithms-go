package binarysearchlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchList(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("Should return -1 when value does not exist (high)", func(t *testing.T) {
		index := Search(haystack, 11)
		assert.Equal(t, index, -1)
	})
	t.Run("Should return -1 when value does not exist (low)", func(t *testing.T) {
		index := Search(haystack, 0)
		assert.Equal(t, index, -1)
	})
	t.Run("Should index when it exist", func(t *testing.T) {
		index := Search(haystack, 4)
		assert.Equal(t, index, 3)
	})
}
