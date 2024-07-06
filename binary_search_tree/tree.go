package binarysearchtree

import "cmp"

type BinarySearchTree[T cmp.Ordered] struct {
	root *Node[T]
}

func (tree *BinarySearchTree[T]) Search(value T) *Node[T] {
	if tree.root == nil {
		return nil
	}

	return tree.root.Search(value)
}

func (tree *BinarySearchTree[T]) Add(value T) *Node[T] {
	if tree.root == nil {
		tree.root = &Node[T]{Value: value}
		return tree.root
	} else {
		return tree.root.Add(value)
	}
}

func (tree *BinarySearchTree[T]) FindMin() *Node[T] {
	if tree.root == nil {
		return nil
	}

	return tree.root.FindMin()
}

func (tree *BinarySearchTree[T]) FindMax() *Node[T] {
	if tree.root == nil {
		return nil
	}

	return tree.root.FindMax()
}

func (tree *BinarySearchTree[T]) Delete(value T) *Node[T] {
	if tree.root == nil {
		return nil
	}

	tree.root = tree.root.Delete(value)
	return tree.root
}
