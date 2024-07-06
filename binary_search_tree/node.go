package binarysearchtree

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (node *Node[T]) Walk(value T) *Node[T] {
	if node.Value >= value && node.Left != nil {
		return node.Left
	}

	if node.Value < value && node.Right != nil {
		return node.Right
	}

	if node.Value == value {
		return node
	}

	return nil
}

func (node *Node[T]) Add(value T) *Node[T] {
	if value <= node.Value && node.Left == nil {
		node.Left = &Node[T]{Value: value}
		return node.Left
	}

	if value > node.Value && node.Right == nil {
		node.Right = &Node[T]{Value: value}
		return node.Right
	}

	return node.Walk(value).Add(value)
}

func (node *Node[T]) Search(value T) *Node[T] {
	if node.Value == value {
		return node
	}

	if next := node.Walk(value); next != nil {
		return next.Search(value)
	}

	return nil
}

func (node *Node[T]) FindMin() *Node[T] {
	if node.Left == nil {
		return node
	}

	return node.Left.FindMin()
}

func (node *Node[T]) FindMax() *Node[T] {
	if node.Right == nil {
		return node
	}

	return node.Right.FindMax()
}

func (node *Node[T]) Delete(value T) *Node[T] {
	if node.Value == value {
		if node.Left != nil && node.Right != nil {
			node.Value = node.Right.FindMin().Value
			node.Right = node.Right.Delete(node.Value)
		} else if node.Left == nil && node.Right == nil {
			node = nil
		} else if node.Left == nil {
			node = node.Right
		} else {
			node = node.Left
		}
	} else if next := node.Walk(value); next != nil {
		next.Delete(value)
	}

	return node
}

func (node *Node[T]) String() string {
	return fmt.Sprintf("Node{Value:%v}", node.Value)
}
