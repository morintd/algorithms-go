package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSearch(t *testing.T) {
	tree := BinaryNode[int]{Value: 1,
		Left: &BinaryNode[int]{Value: 5,
			Left:  &BinaryNode[int]{Value: 6},
			Right: &BinaryNode[int]{Value: 7},
		},
		Right: &BinaryNode[int]{Value: 10,
			Left:  &BinaryNode[int]{Value: 11},
			Right: &BinaryNode[int]{Value: 12},
		},
	}

	t.Run("BreadthFirstSearch", func(t *testing.T) {
		t.Run("Should return true if value is in tree", func(t *testing.T) {
			found := BreadthFirstSearch(&tree, 10)
			assert.Equal(t, true, found)
		})

		t.Run("Should return false if value is not in tree", func(t *testing.T) {
			found := BreadthFirstSearch(&tree, 20)
			assert.Equal(t, false, found)
		})
	})
}
