package model

import "fmt"

type Node[T any] struct {
	id       *int
	Value    T
	Parent   *Node[T]
	Children []*Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		id:       nil,
		Value:    value,
		Parent:   nil,
		Children: nil,
	}
}

func (n *Node[T]) GetId() int {
	return *n.id
}

func (n *Node[T]) SetId(id int) {
	if id <= 0 {
		panic("Node's ID must be positive")
	}
	n.id = &id
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *Node[T]) AddChild(child *Node[T]) {
	child.Parent = n
	n.Children = append(n.Children, child)
}

func (n *Node[T]) RemoveChild(child *Node[T]) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			child.Parent = nil
			break
		}
	}
}

func ParentsIncludeMyself[T any](n *Node[T]) []*Node[T] {
	var parents []*Node[T]
	parents = append(parents, n)
	for p := n.Parent; p != nil; p = p.Parent {
		parents = append(parents, p)
	}
	return parents
}
