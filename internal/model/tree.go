package model

type Tree[T any] struct {
	Root *Node[T]
}

func NewTree[T any](rootLabel Label[T]) *Tree[T] {
	root := NewNode(rootLabel)
	return &Tree[T]{Root: root}
}
