package heap

import "cmp"

type MinHeap[T cmp.Ordered] struct {
	Length int
	data   []T
}

func (heap *MinHeap[T]) Insert(value T) {
	heap.data = append(heap.data, value)
	heap.heapifyUp(heap.Length)
	heap.Length++
}

func (heap *MinHeap[T]) Delete() *T {
	if heap.Length == 0 {
		return nil
	}

	out := heap.data[0]
	heap.Length--

	if heap.Length == 0 {
		heap.data = []T{}
	} else {
		heap.data[0] = heap.data[heap.Length]
		heap.heapifyDown(0)
		heap.data = heap.data[:len(heap.data)-1]
	}

	return &out
}

func (heap *MinHeap[T]) heapifyUp(index int) {
	if index == 0 {
		return
	}

	currentValue := heap.data[index]

	parentIndex := heap.parent(index)
	parentValue := heap.data[parentIndex]

	if parentValue > currentValue {
		heap.data[index] = parentValue
		heap.data[parentIndex] = currentValue

		heap.heapifyUp(parentIndex)
	}
}

func (heap *MinHeap[T]) heapifyDown(index int) {
	leftIndex := heap.leftChild(index)
	rightIndex := heap.rightChild(index)

	if index >= heap.Length || leftIndex >= heap.Length {
		return
	}

	currentValue := heap.data[index]
	leftValue := heap.data[leftIndex]
	rightValue := heap.data[rightIndex]

	if leftValue > rightValue && currentValue > rightValue {
		heap.data[index] = rightValue
		heap.data[rightIndex] = currentValue

		heap.heapifyDown(rightIndex)
	}

	if rightValue > leftValue && currentValue > leftValue {
		heap.data[index] = leftValue
		heap.data[leftIndex] = currentValue

		heap.heapifyDown(leftIndex)
	}
}

func (heap *MinHeap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (heap *MinHeap[T]) leftChild(index int) int {
	return index*2 + 1
}

func (heap *MinHeap[T]) rightChild(index int) int {
	return index*2 + 2
}
