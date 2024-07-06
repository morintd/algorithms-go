package tree

func Compare[T comparable](first *BinaryNode[T], second *BinaryNode[T]) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil || second == nil {
		return false
	}

	if first.Value != second.Value {
		return false
	}

	return Compare(first.Left, second.Left) && Compare(first.Right, second.Right)
}
