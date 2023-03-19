package linkedlist

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Value T

	Next *Node[T]
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
	newNode := Node[T]{Value: val}
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

func (ll *LinkedList[T]) RemoveFirst(val T) {
	if ll.Head.Value == val {
		ll.Head = ll.Head.Next
		return
	}

	prev := ll.Head
	for n := prev.Next; n != nil; n = n.Next {
		if n.Value == val {
			prev.Next = n.Next
			if n == ll.Tail {
				ll.Tail = prev
			}
			return
		}
	}
}

func (ll *LinkedList[T]) Size() int {
	counter := 0
	for n := ll.Head; n != nil; n = n.Next {
		counter++
	}
	return counter
}
