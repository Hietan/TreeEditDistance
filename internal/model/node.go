package model

type Node[T any] struct {
	Value    T
	Parent   *Node[T]
	Children []*Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		Parent:   nil,
		Children: nil,
	}
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
