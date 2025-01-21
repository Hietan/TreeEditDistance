package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func CalcEditDistance[T any](t1, t2 *model.Tree[T]) int {
	t1Size := t1.Size()
	t2Size := t2.Size()

	for i := 1; i <= t1Size; i++ {
		for j := 1; j <= t2Size; j++ {

		}
	}
}
