package lru

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type LRU[K comparable, V comparable] struct {
	capacity int
	length   int

	head *Node[V]
	tail *Node[V]

	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func (lru *LRU[K, V]) Get(key K) *V {
	node, ok := lru.lookup[key]

	if !ok {
		return nil
	}

	lru.detach(node)
	lru.prepend(node)

	return &node.value
}

func (lru *LRU[K, V]) Update(key K, value V) {
	node, ok := lru.lookup[key]

	if ok {
		lru.detach(node)
		node.value = value
	} else {
		if len(lru.lookup) == lru.capacity {
			tailKey := lru.reverseLookup[lru.tail]

			delete(lru.lookup, tailKey)
			delete(lru.reverseLookup, lru.tail)

			lru.detach(lru.tail)
		}

		node = &Node[V]{value: value, prev: nil, next: nil}
		lru.lookup[key] = node
		lru.reverseLookup[node] = key
	}

	lru.prepend(node)
}

func (lru *LRU[K, V]) Length() int {
	return lru.length
}

func (lru *LRU[K, V]) detach(node *Node[V]) {
	lru.length--

	if lru.head == node {
		lru.head = node.next
	}

	if lru.tail == node {
		lru.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	node.next = nil
	node.prev = nil
}

func (lru *LRU[K, V]) prepend(node *Node[V]) {
	lru.length++

	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}

	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}

func NewLRU[K comparable, V comparable](capacity int) *LRU[K, V] {
	lookup := make(map[K]*Node[V], capacity)
	reverseLookup := make(map[*Node[V]]K, capacity)

	return &LRU[K, V]{capacity: capacity, length: 0, head: nil, tail: nil, lookup: lookup, reverseLookup: reverseLookup}
}
