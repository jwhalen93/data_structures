package structures

// Node is the basic building block for many data structures, containing a value and a pointer to the next node.
type Node struct {
	Val  int
	Next *Node
}
