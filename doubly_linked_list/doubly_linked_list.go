package doublylinkedlist

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func (list *DoublyLinkedList[T]) Prepend(item T) {
	node := Node[T]{Value: item}

	list.Length++

	if list.head == nil {
		list.head = &node
		list.tail = &node
		return
	}

	node.Next = list.head
	list.head.Prev = &node
	list.head = &node
}

func (list *DoublyLinkedList[T]) InsertAt(item T, index int) {
	if index > list.Length {
		panic("cannot insert at an index outside of the DoubleLinkedList length")
	}

	if index == list.Length {
		list.Append(item)
		return
	}

	if index == 0 {
		list.Prepend(item)
		return
	}

	list.Length++
	current := list.getAt(index)

	node := Node[T]{Value: item, Next: current, Prev: current.Prev}

	current.Prev.Next = &node
	current.Prev = &node

}

func (list *DoublyLinkedList[T]) Append(item T) {
	list.Length++

	node := Node[T]{Value: item}

	if list.tail == nil {
		list.head = &node
		list.tail = &node
		return
	}

	node.Prev = list.tail
	list.tail = &node
}

func (list *DoublyLinkedList[T]) Remove(item T) *T {
	current := list.head

	for i := 0; i < list.Length; i++ {
		if current.Value == item {
			break
		}

		current = current.Next
	}

	if current == nil {
		return nil
	}

	return list.removeNode(current)
}

func (list *DoublyLinkedList[T]) Get(index int) *T {
	node := list.getAt(index)

	if node == nil {
		return nil
	}

	return &node.Value
}

func (list *DoublyLinkedList[T]) RemoveAt(index int) *T {
	node := list.getAt(index)

	if node == nil {
		return nil
	}

	return list.removeNode(node)
}

func (list *DoublyLinkedList[T]) getAt(index int) *Node[T] {
	current := list.head

	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current
}

func (list *DoublyLinkedList[T]) removeNode(node *Node[T]) *T {
	list.Length--

	if list.Length == 0 {
		out := list.head.Value

		list.head = nil
		list.tail = nil

		return &out
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == list.head {
		list.head = node.Next
	}

	if node == list.tail {
		list.tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil

	return &node.Value
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{Length: 0, head: nil}
}
