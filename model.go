package linkedlist

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Value T

	Previous *Node[T]
	Next     *Node[T]
}

func NewLinkedList[T any](elem T) LinkedList[T] {
	node := &Node[T]{
		Value: elem,
	}

	return LinkedList[T]{
		Head: node,
		Tail: node,
	}
}

func (ll LinkedList[T]) Add(val T) {
	newNode := Node[T]{Previous: ll.Tail}
	ll.Tail = &newNode
}
