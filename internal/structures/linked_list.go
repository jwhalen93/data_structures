package structures

import (
	"errors"
	"fmt"
)

// LinkedList is an implementation of a singly linked list using nodes
type LinkedList struct {
	Head *Node
}

// Append puts a new node at the end of the list
func (l *LinkedList) Append(n Node) {
	head := l.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &n

}

// InsertValue makes it easier to discretely insert numbers directly into the list
func (l *LinkedList) InsertValue(x int) {
	node := Node{Val: x, Next: nil}
	l.Insert(node)
}

// AppendValue makes it easier to discretely append numbers directly into the list
func (l *LinkedList) AppendValue(x int) {
	node := Node{Val: x, Next: nil}
	l.Append(node)
}

// Insert puts a new node at the front of the list
func (l *LinkedList) Insert(n Node) {
	n.Next = l.Head
	l.Head = &n
}

// RemoveLast removes the last element in the list
func (l *LinkedList) RemoveLast() (node *Node, err error) {
	head := l.Head
	if head.Next == nil {
		return nil, errors.New("Cannot removeLast for single node list")
	}
	for head.Next.Next != nil {
		head = head.Next
	}
	removedNode := head.Next
	head.Next = nil

	return removedNode, nil

}

// Remove will remove a node at the given index, using 0 based indexing
func (l *LinkedList) Remove(index int) (n *Node, e error) {
	if index > l.Size()-1 {
		return nil, errors.New("Error removing node, given index is out of range")
	}
	current := l.Head
	prev := current
	currentIndex := 0
	for currentIndex < index {
		prev = current
		current = current.Next
		currentIndex++
	}
	prev.Next = current.Next
	current.Next = nil

	return current, nil

}

// Size returns the count of elements in the list
func (l *LinkedList) Size() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// PrintList iterates through the list, printing every element in order
func (l *LinkedList) PrintList() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func (l *LinkedList) get(index int) (n *Node, e error) {
	listSize := l.Size()
	if index > listSize-1 {
		return nil, errors.New("Error removing node, given index is out of range")
	}
	current := l.Head
	for counter := 0; counter < index; counter++ {
		current = current.Next
	}
	return current, nil
}
