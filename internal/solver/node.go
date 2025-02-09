package solver

import (
	"fmt"
)

type Node[T any] struct {
	value    T
	parent   int
	children []int
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value:    value,
		parent:   -1,
		children: []int{},
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) GetParent() int {
	return n.parent
}

func (n *Node[T]) SetParent(parent int) {
	n.parent = parent
}

func (n *Node[T]) GetChildren() []int {
	return n.children
}

func (n *Node[T]) AddChild(child int) {
	n.children = append(n.children, child)
}

func (n *Node[T]) RemoveChild(child int) {
	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			break
		}
	}
}
