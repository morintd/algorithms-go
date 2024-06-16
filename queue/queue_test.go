package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	t.Run("Should be empty by default", func(t *testing.T) {
		var expected *Node

		queue := Queue{}
		node := queue.Peek()

		assert.Equal(t, node, expected)
		assert.Equal(t, queue.Length, 0)
	})

	t.Run("Should return single element", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue(1)

		node := queue.Peek()
		assert.Equal(t, node.Value, 1)
		assert.Equal(t, queue.Length, 1)
	})

	t.Run("Should return first element", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Enqueue(4)

		dequed := queue.Deque()
		node := queue.Peek()

		assert.Equal(t, dequed.Value, 1)
		assert.Equal(t, node.Value, 2)
		assert.Equal(t, queue.Length, 3)
	})
}
