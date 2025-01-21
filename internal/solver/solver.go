package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func CalcEditDistance[T any](t1, t2 *model.Tree[T]) int {
	fmt.Println(t1.Size(), t2.Size())
	return 0
}
