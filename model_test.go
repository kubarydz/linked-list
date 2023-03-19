package linkedlist

import "testing"

func TestContains(t *testing.T) {
	t.Run("contains added object", func(t *testing.T) {
		list := NewLinkedList(1)

		list.Add(4)
		list.Add(2)
		list.Add(0)
		list.Add(6)

		if !list.Contains(2) {
			t.Errorf("added item not found in list\n")
		}
	})

}

func TestRemove(t *testing.T) {
	t.Run("remove existing object", func(t *testing.T) {
		list := NewLinkedList(1)

		list.Add(4)
		list.Add(2)
		list.Add(0)
		list.Add(6)

		list.RemoveFirst(2)
		if list.Contains(2) {
			t.Errorf("didn't remove the object\n")
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("calculate size", func(t *testing.T) {
		list := NewLinkedList(1)

		list.Add(2)
		list.Add(11)

		if list.Size() != 3 {
			t.Errorf("size calculation incorrect\n")
		}

	})
}
