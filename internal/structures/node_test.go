package structures

import (
	"testing"
)

func TestNode(t *testing.T) {
	node1 := Node{1, nil}
	node2 := Node{2, nil}

	node1.Next = &node2

	if node1.Next.Val != 2 {
		t.Errorf("Next value was incorrect, got %d, should be %d.", node1.Next.Val, 2)
	}
}
