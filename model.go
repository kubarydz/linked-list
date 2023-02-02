package linkedlist

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Previous *Node[T]
	Next     *Node[T]
}

func (ll LinkedList[T]) Add(val T) {
	newNode := Node[T]{Previous: ll.Tail}
	ll.Tail = &newNode
}
