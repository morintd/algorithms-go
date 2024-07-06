package tree

func BreadthFirstSearch[T comparable](head *BinaryNode[T], value T) bool {
	queue := []*BinaryNode[T]{head}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Value == value {
			return true
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return false
}
