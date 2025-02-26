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

func Cost[T comparable](indBefore int, indAfter int, treeBefore Tree[T], treeAfter Tree[T]) int {

	if indBefore == EmptyIndex && indAfter == EmptyIndex {
		return 0
	} else if indBefore == EmptyIndex || indAfter == EmptyIndex {
		return 1
	}

	if treeBefore.GetNodes()[indBefore].GetValue() == treeAfter.GetNodes()[indAfter].GetValue() {
		return 0
	} else {
		return 1
	}
}

func (t *Tree[T]) GetPathToRoot(targetIndex int) *[]int {
	path := []int{}
	now := targetIndex
	for {
		now = t.GetNodes()[now].GetParent()
		if now == EmptyIndex {
			break
		}
		path = append(path, now)
	}
	return &path
}

func (t *Tree[T]) GetPathToRootIncludeMyself(targetIndex int) []int {
	path := *t.GetPathToRoot(targetIndex)
	return append([]int{targetIndex}[:], path[:]...)
}

func (t *Tree[T]) GetChildOnPath(parent int, descendant int) int {
	path := t.GetPathToRootIncludeMyself(descendant)

	for i := len(path) - 1; i > 0; i-- {
		if path[i] == parent {
			return path[i-1]
		}
	}

	return -1
}
