package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeTraversal(t *testing.T) {
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

	t.Run("PreOrder", func(t *testing.T) {
		t.Run("Should return PreOrder traversed", func(t *testing.T) {
			traversed := PreOrderSearch(&tree)
			assert.Equal(t, []int{1, 5, 6, 7, 10, 11, 12}, traversed)
		})
	})

	t.Run("InOrder", func(t *testing.T) {
		t.Run("Should return InOrder traversed", func(t *testing.T) {
			traversed := InOrderSearch(&tree)
			assert.Equal(t, []int{5, 6, 7, 1, 10, 11, 12}, traversed)
		})
	})

	t.Run("PostOrder", func(t *testing.T) {
		t.Run("Should return PostOrder traversed", func(t *testing.T) {
			traversed := PostOrderSearch(&tree)
			assert.Equal(t, []int{5, 6, 7, 10, 11, 12, 1}, traversed)
		})
	})
}
