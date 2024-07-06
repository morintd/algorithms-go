package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		t.Run("Should set value as head/tail if empty", func(t *testing.T) {
			list := DoublyLinkedList[int]{}
			list.Prepend(1)

			assert.Equal(t, &Node[int]{Value: 1}, list.head)
			assert.Equal(t, &Node[int]{Value: 1}, list.tail)
		})

		t.Run("Should increase length", func(t *testing.T) {
			list := DoublyLinkedList[int]{}
			list.Prepend(1)

			assert.Equal(t, 1, list.Length)
		})

		t.Run("Should set value as head linked to previous head otherwise", func(t *testing.T) {
			head := Node[int]{Value: 1, Prev: nil, Next: nil}

			list := DoublyLinkedList[int]{head: &head, tail: &head, Length: 0}
			list.Prepend(2)

			assert.Equal(t, &Node[int]{Value: 2, Prev: nil, Next: &head}, list.head)
			assert.Equal(t, &head, list.head.Next)
		})
	})

	t.Run("InsertAt", func(t *testing.T) {
		initial := Node[int]{Value: 1}
		list := DoublyLinkedList[int]{head: &initial, tail: &initial, Length: 1}
		list.InsertAt(2, 1)

		inserted := Node[int]{Value: 2, Prev: &initial}
		initial.Next = &inserted

		t.Run("Should insert at index", func(t *testing.T) {
			assert.Equal(t, &initial, list.head)
		})

		t.Run("Should increment length", func(t *testing.T) {
			assert.Equal(t, 2, list.Length)
		})
	})

	t.Run("Append", func(t *testing.T) {
		t.Run("Should set value as head/tail if empty", func(t *testing.T) {
			list := DoublyLinkedList[int]{}
			list.Append(1)

			assert.Equal(t, &Node[int]{Value: 1}, list.head)
			assert.Equal(t, &Node[int]{Value: 1}, list.tail)
		})

		t.Run("Should increase length", func(t *testing.T) {
			list := DoublyLinkedList[int]{}
			list.Prepend(1)

			assert.Equal(t, 1, list.Length)
		})

		t.Run("Should set value as tail linked to previous tail otherwise", func(t *testing.T) {
			initial := Node[int]{Value: 1, Prev: nil, Next: nil}

			list := DoublyLinkedList[int]{head: &initial, tail: &initial, Length: 0}
			list.Append(2)

			assert.Equal(t, &Node[int]{Value: 2, Prev: &initial, Next: nil}, list.tail)
			assert.Equal(t, &initial, list.tail.Prev)
		})
	})

	t.Run("Remove", func(t *testing.T) {

		t.Run("single element", func(t *testing.T) {
			initial := Node[int]{Value: 1}
			list := DoublyLinkedList[int]{head: &initial, tail: &initial, Length: 1}
			removed := list.Remove(1)

			t.Run("Should return removed", func(t *testing.T) {
				assert.Equal(t, 1, *removed)
			})

			t.Run("Should remove element", func(t *testing.T) {
				var expected *Node[int]
				assert.Equal(t, expected, list.head)
				assert.Equal(t, expected, list.tail)
			})

			t.Run("Should decrement length", func(t *testing.T) {
				assert.Equal(t, 0, list.Length)
			})
		})

		t.Run("in the middle", func(t *testing.T) {
			head := Node[int]{Value: 1}
			node := Node[int]{Value: 2}
			tail := Node[int]{Value: 3}

			head.Next = &node
			node.Prev = &head
			node.Next = &tail
			tail.Prev = &node

			list := DoublyLinkedList[int]{head: &head, tail: &tail, Length: 3}
			removed := list.Remove(2)

			t.Run("Should return removed", func(t *testing.T) {
				assert.Equal(t, 2, *removed)
			})

			t.Run("Should remove element", func(t *testing.T) {
				head.Next = &tail
				tail.Prev = &head

				assert.Equal(t, &head, list.head)
				assert.Equal(t, &tail, list.tail)
			})

			t.Run("Should decrement length", func(t *testing.T) {
				assert.Equal(t, 2, list.Length)
			})
		})
	})

	t.Run("Get", func(t *testing.T) {
		head := Node[int]{Value: 1}
		node := Node[int]{Value: 2}
		tail := Node[int]{Value: 3}

		head.Next = &node
		node.Prev = &head
		node.Next = &tail
		tail.Prev = &node

		list := DoublyLinkedList[int]{head: &head, tail: &tail, Length: 3}

		t.Run("Should return value of node at index", func(t *testing.T) {
			result := list.Get(1)
			assert.Equal(t, *result, 2)
		})
	})

	t.Run("RemoveAt", func(t *testing.T) {

		t.Run("single element", func(t *testing.T) {
			initial := Node[int]{Value: 1}
			list := DoublyLinkedList[int]{head: &initial, tail: &initial, Length: 1}
			removed := list.RemoveAt(0)

			t.Run("Should return removed", func(t *testing.T) {
				assert.Equal(t, 1, *removed)
			})

			t.Run("Should remove element", func(t *testing.T) {
				var expected *Node[int]
				assert.Equal(t, expected, list.head)
				assert.Equal(t, expected, list.tail)
			})

			t.Run("Should decrement length", func(t *testing.T) {
				assert.Equal(t, 0, list.Length)
			})
		})

		t.Run("in the middle", func(t *testing.T) {
			head := Node[int]{Value: 1}
			node := Node[int]{Value: 2}
			tail := Node[int]{Value: 3}

			head.Next = &node
			node.Prev = &head
			node.Next = &tail
			tail.Prev = &node

			list := DoublyLinkedList[int]{head: &head, tail: &tail, Length: 3}
			removed := list.RemoveAt(1)

			t.Run("Should return removed", func(t *testing.T) {
				assert.Equal(t, 2, *removed)
			})

			t.Run("Should remove element", func(t *testing.T) {
				head.Next = &tail
				tail.Prev = &head

				assert.Equal(t, &head, list.head)
				assert.Equal(t, &tail, list.tail)
			})

			t.Run("Should decrement length", func(t *testing.T) {
				assert.Equal(t, 2, list.Length)
			})
		})
	})
}
