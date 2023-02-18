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
			t.Errorf("added item not found in list")
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

		list.Remove(2)
		if list.Contains(2) {
			t.Errorf("didn't remove the object")
		}
	})
}
