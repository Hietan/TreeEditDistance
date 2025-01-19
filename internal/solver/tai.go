package solver

import "github.com/Hietan/TreeEditDistance/internal/model"

type Tai[T any] struct{}

func (s *Tai[T]) CalcEditDistance(tree1, tree2 *model.Tree[T]) int {
	return 0
}
