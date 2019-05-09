package structures

import "testing"

func TestLinkedList(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	if linkedList.Head.Val != 1 {
		t.Errorf("Error creating list, got value %d, expected value %d", linkedList.Head.Val, 1)
	}
}
func TestAppend(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	newNode := Node{Val: 2, Next: nil}

	linkedList.Append(newNode)

	if linkedList.Head.Next.Val != 2 {
		t.Errorf("Error inserting value into list, got value %d, expected value %d", linkedList.Head.Next.Val, 2)
	}

	if linkedList.Size() != 2 {
		t.Errorf("Error getting size of list, got value %d, expected value %d", linkedList.Size(), 2)
	}
}
func TestRemoveLast(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	newNode := Node{Val: 2, Next: nil}

	linkedList.Append(newNode)

	removedNode, err := linkedList.RemoveLast()

	if err != nil {
		t.Errorf("Error while removing last node of the list, %s", err)
	}

	if removedNode.Val != 2 {
		t.Errorf("Error when removing last node, invalid value, got value %d, expected value %d", removedNode.Val, 2)
	}

	if linkedList.Head.Next != nil {
		t.Errorf("Error when removing last node, last node not actually removed, %+v ", linkedList.Head.Next)
	}

	removedNode, err = linkedList.RemoveLast()

	if err == nil {
		t.Errorf("Error while removing last, expecting error, didn't recieve it")
	}

	if linkedList.Size() != 1 {
		t.Errorf("Error when getting size after removal, got value %d, expected %d", linkedList.Size(), 1)
	}

}

func TestInsert(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	insertNode := Node{Val: 0, Next: nil}

	linkedList.Insert(insertNode)

	if linkedList.Head.Val != 0 {
		t.Errorf("Error when inserting node, got value %d, expected %d", linkedList.Head.Val, 0)
	}
}

func TestRemove(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	linkedList.AppendValue(2)
	linkedList.AppendValue(3)

	removedNode, _ := linkedList.Remove(1)

	if linkedList.Head.Next.Val != 3 {
		t.Errorf("Error removing node at specific index, got %d, expected %d", linkedList.Head.Next.Val, 3)
	}

	if removedNode.Val != 2 {
		t.Errorf("Error removed node value, got %d, expected %d", removedNode.Val, 2)
	}

}

func TestGet(t *testing.T) {
	head := Node{Val: 1, Next: nil}
	linkedList := LinkedList{Head: &head}

	linkedList.AppendValue(2)
	linkedList.AppendValue(3)

	getNode, err := linkedList.get(1)

	if err != nil {
		t.Errorf("Error getting node at specific index: %s", err)
	}

	if getNode.Val != 2 {
		t.Errorf("Error with retrieved node, got value %d, expected %d", getNode.Val, 2)
	}

	if linkedList.Size() != 3 {
		t.Errorf("Error with getting node, size incorrect, expected %d, got %d", linkedList.Size(), 3)
	}
}
