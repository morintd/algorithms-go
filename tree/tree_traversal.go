package tree

func walkPreOrder(current *BinaryNode[int], path []int) []int {
	if current == nil {
		return path
	}

	path = append(path, current.Value)

	path = walkPreOrder(current.Left, path)
	path = walkPreOrder(current.Right, path)

	return path
}

func walkInOrder(current *BinaryNode[int], path []int) []int {
	if current == nil {
		return path
	}

	path = walkPreOrder(current.Left, path)
	path = append(path, current.Value)
	path = walkPreOrder(current.Right, path)

	return path
}

func walkPostOrder(current *BinaryNode[int], path []int) []int {
	if current == nil {
		return path
	}

	path = walkPreOrder(current.Left, path)
	path = walkPreOrder(current.Right, path)

	path = append(path, current.Value)

	return path
}

func PreOrderSearch(root *BinaryNode[int]) []int {
	path := []int{}
	path = walkPreOrder(root, path)

	return path
}

func InOrderSearch(root *BinaryNode[int]) []int {
	path := []int{}
	path = walkInOrder(root, path)

	return path
}

func PostOrderSearch(root *BinaryNode[int]) []int {
	path := []int{}
	path = walkPostOrder(root, path)

	return path
}
