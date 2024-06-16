package queue

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Length int
	head   *Node
	tail   *Node
}

func (queue *Queue) Enqueue(value int) *Node {
	node := &Node{
		Value: value,
	}

	if queue.tail == nil {
		queue.head = node
	} else {
		queue.tail.Next = node
	}

	queue.tail = node
	queue.Length++

	return node
}

func (queue *Queue) Deque() *Node {
	if queue.head == nil {
		return nil
	}

	node := queue.head
	queue.head = queue.head.Next

	if queue.head == nil {
		queue.tail = nil
	}

	queue.Length--
	return node
}

func (queue *Queue) Peek() *Node {
	return queue.head
}
