package tree

type BinaryNode[T comparable] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}
