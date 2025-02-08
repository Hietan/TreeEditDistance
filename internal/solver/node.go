package solver

import "fmt"

type Node[T any] struct {
	id       int
	value    T
	parent   *int
	children []int
}

func NewNode[T any](id int, value T) *Node[T] {
	return &Node[T]{
		id:       id,
		value:    value,
		parent:   nil,
		children: nil,
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.value)
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

func (n *Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) GetParent() *Node[T] {
	return n.parent
}

func (n *Node[T]) SetParent(parent *Node[T]) {
	n.parent = parent
}

func (n *Node[T]) GetChildren() []*Node[T] {
	return n.children
}

func (n *Node[T]) AddChild(child *Node[T]) {
	child.SetParent(n)
	n.children = append(n.children, child)
}

func (n *Node[T]) RemoveChild(child *Node[T]) {
	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			child.SetParent(nil)
			break
		}
	}
}

func parentsIncludeMyself[T any](n *Node[T]) []*Node[T] {
	var parents []*Node[T]
	parents = append(parents, n)
	for p := n.GetParent(); p != nil; p = p.GetParent() {
		parents = append(parents, p)
	}
	return parents
}
