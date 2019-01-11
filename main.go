package main

import (
	"fmt"
)

type Tree struct {
	Parent   *Tree
	Children map[string]*Tree
	Payload  bool // Your data here
}

func NewTree(parent *Tree, path []string, payload bool) *Tree {
	if parent == nil {
		parent = &Tree{nil, map[string]*Tree{}, false}
	}
	if len(path) == 0 {
		parent.Payload = payload
		return parent
	}

	child := parent.Children[path[0]]
	if child == nil {
		child = &Tree{parent, map[string]*Tree{}, false}
		parent.Children[path[0]] = child
	}
	return NewTree(child, path[1:], payload)
}

func main() {
	root := NewTree(nil, nil, false)
	NewTree(root, []string{"A", "B", "C"}, true)
	NewTree(root, []string{"A", "B", "D"}, true)
	fmt.Print(root.Children["A"].Children["B"].Children["D"].Payload)
}
