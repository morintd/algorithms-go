package binarysearchtree

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) Search(value int) *Node {
	if tree.root == nil {
		return nil
	}

	return tree.root.Search(value)
}

func (tree *BinarySearchTree) Add(value int) *Node {
	if tree.root == nil {
		tree.root = &Node{Value: value}
		return tree.root
	} else {
		return tree.root.Add(value)
	}
}

func (tree *BinarySearchTree) FindMin() *Node {
	if tree.root == nil {
		return nil
	}

	return tree.root.FindMin()
}

func (tree *BinarySearchTree) FindMax() *Node {
	if tree.root == nil {
		return nil
	}

	return tree.root.FindMax()
}

func (tree *BinarySearchTree) Delete(value int) *Node {
	if tree.root == nil {
		return nil
	}

	tree.root = tree.root.Delete(value)
	return tree.root
}
