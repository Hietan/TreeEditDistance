package model

import (
	"fmt"
	"strings"
)

type Tree[T any] struct {
	root *Node[T]
}

func NewTree[T any](rootValue T) *Tree[T] {
	root := NewNode(rootValue)
	return &Tree[T]{root: root}
}

func (t *Tree[T]) String() string {
	if t.root == nil {
		return "<empty tree>"
	}

	var sb strings.Builder
	var printTree func(node *Node[T], prefix string, isRoot bool, isLast bool)
	printTree = func(node *Node[T], prefix string, isRoot bool, isLast bool) {
		if node == nil {
			return
		}

		if !isRoot {
			sb.WriteString(prefix)
			if isLast {
				sb.WriteString("└── ")
				prefix += "    "
			} else {
				sb.WriteString("├── ")
				prefix += "│   "
			}
		}

		sb.WriteString(fmt.Sprintf("%v\n", node.GetValue()))

		for i, child := range node.GetChildren() {
			printTree(child, prefix, false, i == len(node.GetChildren())-1)
		}
	}

	printTree(t.root, "", true, true)
	return sb.String()
}

func (t *Tree[T]) GetRoot() *Node[T] {
	return t.root
}

func subtreeSize[T any](n *Node[T]) int {
	size := 1
	if n.GetChildren() == nil {
		return 1
	}
	for _, child := range n.GetChildren() {
		size += subtreeSize(child)
	}
	return size
}

func (t *Tree[T]) Size() int {
	return subtreeSize(t.root)
}
