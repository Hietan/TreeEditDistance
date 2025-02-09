package solver

import (
	"github.com/Hietan/TreeEditDistance/internal/model"
)

const EmptyIndex = -1

type Tree[T any] struct {
	nodes  []*Node[T]
	length int
}

func (t *Tree[T]) GetNodes() []*Node[T] {
	return t.nodes
}

func (t *Tree[T]) GetLength() int {
	return t.length
}

func traverse[T any](nodes *[]*Node[T], now *model.Node[T], index int, parent int) int {
	node := NewNode(now.GetValue())
	node.SetParent(parent)

	*nodes = append(*nodes, node)

	next := index + 1
	for _, child := range now.GetChildren() {
		node.AddChild(next)
		next = traverse(nodes, child, next, index)
	}

	return next
}

func NewTree[T any](t *model.Tree[T]) *Tree[T] {
	nodes := []*Node[T]{}
	length := traverse(&nodes, t.GetRoot(), 0, EmptyIndex)

	return &Tree[T]{
		nodes:  nodes,
		length: length,
	}
}

func Cost(beforeInd int, afterInd int) int {
	return 1
}

func (t *Tree[T]) GetParents(targetIndex int) *[]int {
	parents := []int{}
	now := targetIndex
	for {
		parents = append(parents, now)
		now = t.GetNodes()[now].GetParent()
		if now == EmptyIndex {
			break
		}
	}
	return &parents
}
