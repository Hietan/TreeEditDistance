package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func addSubtreeToArray[T any](arr []*model.Node[T], n *model.Node[T]) []*model.Node[T] {
	arr = append(arr, n)
	children := n.GetChildren()
	if children != nil {
		for _, child := range children {
			arr = addSubtreeToArray(arr, child)
		}
	}
	return arr
}

func makeNodeArray[T any](t *model.Tree[T]) []*model.Node[T] {
	return addSubtreeToArray([]*model.Node[T]{}, t.GetRoot())
}

func CalcEditDistance[T any](t1, t2 *model.Tree[T]) int {
	nodeArrayT1 := makeNodeArray(t1)
	nodeArrayT2 := makeNodeArray(t2)

	fmt.Println(len(nodeArrayT1), len(nodeArrayT2))

	sizeT1 := t1.Size()
	sizeT2 := t2.Size()

	for i := 1; i <= sizeT1; i++ {
		for j := 1; j <= sizeT2; j++ {

		}
	}

	return 0
}
