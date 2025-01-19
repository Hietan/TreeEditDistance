package solver

import "github.com/Hietan/TreeEditDistance/internal/model"

type Solver[T any] interface {
	CalcEditDistance(tree1, tree2 *model.Tree[T]) int
}
