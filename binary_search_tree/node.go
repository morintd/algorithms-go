package binarysearchtree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) Walk(value int) *Node {
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

func (node *Node) Add(value int) *Node {
	if value <= node.Value && node.Left == nil {
		node.Left = &Node{Value: value}
		return node.Left
	}

	if value > node.Value && node.Right == nil {
		node.Right = &Node{Value: value}
		return node.Right
	}

	return node.Walk(value).Add(value)
}

func (node *Node) Search(value int) *Node {
	if node.Value == value {
		return node
	}

	if next := node.Walk(value); next != nil {
		return next.Search(value)
	}

	return nil
}

func (node *Node) FindMin() *Node {
	if node.Left == nil {
		return node
	}

	return node.Left.FindMin()
}

func (node *Node) FindMax() *Node {
	if node.Right == nil {
		return node
	}

	return node.Right.FindMax()
}

func (node *Node) Delete(value int) *Node {
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

func (node *Node) String() string {
	return fmt.Sprintf("Node{Value:%d}", node.Value)
}
