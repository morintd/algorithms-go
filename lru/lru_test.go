package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Run("LRU", func(t *testing.T) {
		t.Run("Update", func(t *testing.T) {
			t.Run("When element does not exist ", func(t *testing.T) {
				lru := NewLRU[string, int](10)
				lru.Update("first", 1)

				t.Run("Should create element ", func(t *testing.T) {
					assert.Equal(t, &Node[int]{value: 1, prev: nil, next: nil}, lru.head)
				})

				t.Run("Should increase length ", func(t *testing.T) {
					assert.Equal(t, 1, lru.length)
				})
			})

			t.Run("When element already exists ", func(t *testing.T) {
				lru := NewLRU[string, int](10)
				lru.Update("first", 1)
				lru.Update("first", 10)

				t.Run("Should update element ", func(t *testing.T) {
					assert.Equal(t, &Node[int]{value: 10, prev: nil, next: nil}, lru.head)
				})

				t.Run("Should keep length ", func(t *testing.T) {
					assert.Equal(t, 1, lru.length)
				})
			})
		})

		t.Run("Get", func(t *testing.T) {
			lru := NewLRU[string, int](10)
			lru.Update("first", 1)
			lru.Update("second", 2)
			lru.Update("third", 3)

			v := lru.Get("second")

			t.Run("Should return value", func(t *testing.T) {
				assert.Equal(t, 2, *v)
			})

			t.Run("Should move node to head", func(t *testing.T) {
				assert.Equal(t, 2, lru.head.value)
			})
		})
	})
}
