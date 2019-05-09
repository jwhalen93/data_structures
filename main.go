package main

import (
	"github.com/jwhalen93/data_structures/internal/structures"
)

func main() {
	node := &structures.Node{Val: 1, Next: nil}
	nextNode := &structures.Node{Val: 2, Next: nil}

	list := structures.LinkedList{Head: node}

	list.append(nextNode)
}
