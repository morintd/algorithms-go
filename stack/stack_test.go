package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	t.Run("Should be empty by default", func(t *testing.T) {
		var expected *Node

		stack := Stack{}
		node := stack.Peek()

		assert.Equal(t, node, expected)
		assert.Equal(t, stack.Length, 0)
	})

	t.Run("Should return single element", func(t *testing.T) {
		stack := Stack{}
		stack.Enqueue(1)

		node := stack.Peek()
		assert.Equal(t, node.Value, 1)
		assert.Equal(t, stack.Length, 1)
	})

	t.Run("Should return last element", func(t *testing.T) {
		stack := Stack{}
		stack.Enqueue(1)
		stack.Enqueue(2)
		stack.Enqueue(3)
		stack.Enqueue(4)

		dequed := stack.Deque()
		node := stack.Peek()

		assert.Equal(t, dequed.Value, 4)
		assert.Equal(t, node.Value, 3)
		assert.Equal(t, stack.Length, 3)
	})
}
