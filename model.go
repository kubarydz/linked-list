package linkedlist

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Value T

	Previous *Node[T]
	Next     *Node[T]
}

func NewLinkedList[T comparable](elem T) LinkedList[T] {
	node := &Node[T]{
		Value: elem,
	}

	return LinkedList[T]{
		Head: node,
		Tail: node,
	}
}

func (ll *LinkedList[T]) Add(val T) {
	newNode := Node[T]{Value: val, Previous: ll.Tail}
	ll.Tail.Next = &newNode
	ll.Tail = &newNode
}

func (ll *LinkedList[T]) Contains(val T) bool {
	for n := ll.Head; n != nil; {
		if n.Value == val {
			return true
		}
		n = n.Next
	}
	return false
}

func (ll *LinkedList[T]) Remove(val T) {
	for n := ll.Head; n != nil; {
		if n.Value == val {
			prev := n.Previous
			prev.Next = n.Next
			return
		}
		n = n.Next
	}
}
