package queue

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Length int
	head   *Node
	tail   *Node
}

func (stack *Stack) Enqueue(value int) *Node {
	node := &Node{
		Value: value,
		Next:  stack.head,
	}

	if stack.tail == nil {
		stack.tail = node
	}

	stack.head = node
	stack.Length++

	return node
}

func (stack *Stack) Deque() *Node {
	if stack.head == nil {
		return nil
	}

	node := stack.head
	stack.head = stack.head.Next

	if stack.head == nil {
		stack.tail = nil
	}

	stack.Length--
	return node
}

func (stack *Stack) Peek() *Node {
	return stack.head
}
