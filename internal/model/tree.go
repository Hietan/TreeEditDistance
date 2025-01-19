package model

import (
	"fmt"
	"strings"
)

type Tree[T any] struct {
	Root *Node[T]
}

func NewTree[T any](rootValue T) *Tree[T] {
	root := NewNode(rootValue)
	return &Tree[T]{Root: root}
}

func (t *Tree[T]) String() string {
	if t.Root == nil {
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

		sb.WriteString(fmt.Sprintf("%v\n", node.Value))

		for i, child := range node.Children {
			printTree(child, prefix, false, i == len(node.Children)-1)
		}
	}

	printTree(t.Root, "", true, true)
	return sb.String()
}
