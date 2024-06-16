package binarysearchtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateTestTree() *BinarySearchTree {
	return &BinarySearchTree{
		&Node{
			Value: 50,
			Left: &Node{
				Value: 40,
				Left: &Node{
					Right: nil,
					Left:  nil,
					Value: 35,
				},
				Right: &Node{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
			},
			Right: &Node{
				Value: 60,
				Left: &Node{
					Value: 55,
					Right: nil,
					Left:  nil,
				},
				Right: &Node{
					Value: 65,
					Right: nil,
					Left:  nil,
				},
			},
		},
	}
}

func TestTreeSearch(t *testing.T) {

	t.Run("Should return Node with value", func(t *testing.T) {
		tree := generateTestTree()
		values := []int{35, 65}

		for i, expected := range values {
			t.Run(fmt.Sprintf("[%d] Should return Node with value %d", i, expected), func(t *testing.T) {
				node := tree.Search(expected)
				assert.Equal(t, node.Value, expected)
			})
		}
	})

	t.Run("Should return nil when tree has no root", func(t *testing.T) {
		tree := BinarySearchTree{}
		var expected *Node

		node := tree.Search(0)
		assert.Equal(t, node, expected)
	})

	t.Run("Should return nil when value does not exist", func(t *testing.T) {
		tree := generateTestTree()
		var expected *Node

		node := tree.Search(47)
		assert.Equal(t, node, expected)
	})
}

func TestTreeAdd(t *testing.T) {
	t.Run("Should set root when root is nil", func(t *testing.T) {
		tree := &BinarySearchTree{}
		tree.Add(5)

		assert.Equal(t, tree.root.Value, 5)
	})

	t.Run("Should add at the end of the tree (right)", func(t *testing.T) {
		tree := generateTestTree()
		tree.Add(47)

		assert.Equal(t, tree.root.Left.Right.Right.Value, 47)
	})

	t.Run("Should add at the end of the tree (left)", func(t *testing.T) {
		tree := generateTestTree()
		tree.Add(53)

		assert.Equal(t, tree.root.Right.Left.Left.Value, 53)
	})
}

func TestTreeFindMin(t *testing.T) {
	t.Run("Should return min", func(t *testing.T) {
		tree := generateTestTree()
		node := tree.FindMin()

		assert.Equal(t, node.Value, 35)
	})

	t.Run("Should return nil if root is nil", func(t *testing.T) {
		tree := &BinarySearchTree{}
		var expected *Node

		node := tree.FindMin()
		assert.Equal(t, node, expected)
	})
}

func TestTreeFindMax(t *testing.T) {

	t.Run("Should return max", func(t *testing.T) {
		tree := generateTestTree()
		node := tree.FindMax()

		assert.Equal(t, node.Value, 65)
	})

	t.Run("Should return nil if root is nil", func(t *testing.T) {
		tree := &BinarySearchTree{}
		var expected *Node

		node := tree.FindMax()
		assert.Equal(t, node, expected)
	})
}

func TestTreeRemove(t *testing.T) {

	t.Run("Should return nil if root is nil", func(t *testing.T) {
		tree := &BinarySearchTree{}
		var expected *Node

		node := tree.Delete(50)
		assert.Equal(t, node, expected)
	})

	t.Run("Should remove root if single node", func(t *testing.T) {
		tree := &BinarySearchTree{&Node{
			Value: 50,
			Right: nil,
			Left:  nil,
		}}
		var expected *Node

		tree.Delete(50)
		assert.Equal(t, tree.root, expected)
	})

	t.Run("Should replace deleted nested node", func(t *testing.T) {
		tree := generateTestTree()

		tree.Delete(60)
		assert.Equal(t, tree.root.Right.Value, 65)
		assert.Equal(t, tree.root.Right.Left.Value, 55)
	})

	t.Run("Should replace deleted root", func(t *testing.T) {
		tree := generateTestTree()

		tree.Delete(50)

		assert.Equal(t, tree.root.Value, 55)

		assert.Equal(t, tree.root.Left.Value, 40)
		assert.Equal(t, tree.root.Right.Value, 60)

		assert.Equal(t, tree.root.Left.Left.Value, 35)
		assert.Equal(t, tree.root.Left.Right.Value, 45)

		assert.Equal(t, tree.root.Right.Right.Value, 65)
	})
}
