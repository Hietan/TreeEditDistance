package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func CalcEditDistance[T any](t1, t2 *model.Tree[T]) int {
	tree1 := NewTree(t1)
	//tree2 := NewTree(t2)

	//fmt.Println(tree1.GetLength(), tree2.GetLength())

	//for _, node := range tree1.GetNodes() {
	//	fmt.Println(node.GetValue(), node.GetParent(), node.GetChildren())
	//}

	fmt.Println(tree1.GetParents(4))

	return 0
}
