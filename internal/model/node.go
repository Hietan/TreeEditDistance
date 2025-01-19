package model

type Node[T any] struct {
	Label    Label[T]
	Parent   *Node[T]
	Children []*Node[T]
}

func NewNodeFromLabel[T any](label Label[T]) *Node[T] {
	return &Node[T]{
		Label:    label,
		Parent:   nil,
		Children: nil,
	}
}

func NewNodeFromValue[T any](value T) *Node[T] {
	return &Node[T]{
		Label:    *NewLabel(value),
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
