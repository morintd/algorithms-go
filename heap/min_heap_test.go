package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Run("MinHeap", func(t *testing.T) {
		t.Run("Insert", func(t *testing.T) {
			heap := MinHeap[int]{data: []int{0, 1, 2}, Length: 3}
			heap.Insert(3)

			t.Run("Should add element", func(t *testing.T) {
				assert.Equal(t, []int{0, 1, 2, 3}, heap.data)
			})

			t.Run("Should increase length", func(t *testing.T) {
				assert.Equal(t, 4, heap.Length)
			})
		})

		t.Run("Delete", func(t *testing.T) {
			heap := MinHeap[int]{data: []int{0, 1, 2}, Length: 3}
			heap.Delete()

			t.Run("Should delete element", func(t *testing.T) {
				assert.Equal(t, []int{1, 2}, heap.data)
			})

			t.Run("Should decrease length", func(t *testing.T) {
				assert.Equal(t, 2, heap.Length)
			})
		})

		t.Run("Delete (last element)", func(t *testing.T) {
			heap := MinHeap[int]{data: []int{0}, Length: 1}
			heap.Delete()

			t.Run("Should delete element", func(t *testing.T) {
				assert.Equal(t, []int{}, heap.data)
			})

			t.Run("Should decrease length", func(t *testing.T) {
				assert.Equal(t, heap.Length, 0)
			})
		})
	})
}
