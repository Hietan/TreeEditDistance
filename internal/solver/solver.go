package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func calcE[T comparable](tBef Tree[T], tAft Tree[T]) [][][][][][]int {
	lenBef := tBef.GetLength()
	lenAft := tAft.GetLength()

	e := make([][][][][][]int, lenBef)
	for s := range e {
		e[s] = make([][][][][]int, lenBef)
		for u := range e[s] {
			e[s][u] = make([][][][]int, lenBef)
			for i := range e[s][u] {
				e[s][u][i] = make([][][]int, lenAft)
				for t := range e[s][u][i] {
					e[s][u][i][t] = make([][]int, lenAft)
					for v := range e[s][u][i][t] {
						e[s][u][i][t][v] = make([]int, lenAft)
					}
				}
			}
		}
	}

	for i := 0; i < lenBef; i++ {
		for j := 0; j < lenAft; j++ {
			for _, u := range tBef.GetPathToRootIncludeMyself(i) {
				for _, s := range tBef.GetPathToRootIncludeMyself(u) {
					for _, v := range tAft.GetPathToRootIncludeMyself(j) {
						for _, t := range tAft.GetPathToRootIncludeMyself(v) {
							if (s == u && u == i) && (t == v && v == j) {
								e[s][u][i][t][v][j] = Cost(i, j, tBef, tAft)
							} else if (s == u && u == i) || (t < v && v == j) {
								e[s][u][i][t][v][j] = e[s][u][i][t][tAft.GetNodes()[j].GetParent()][j-1] + Cost(EmptyIndex, j, tBef, tAft)
							} else if (s < u && u == i) || (t == v && v == j) {
								e[s][u][i][t][v][j] = e[s][tBef.GetNodes()[i].GetParent()][i-1][t][v][j] + Cost(i, EmptyIndex, tBef, tAft)
							} else {
								x := tBef.GetChildOnPath(u, i)
								y := tAft.GetChildOnPath(v, j)
								e[s][u][i][t][v][j] = min(e[s][x][i][t][v][j], e[s][u][i][t][y][j], e[s][u][x-1][t][v][y-1]+e[x][x][i][y][y][j])
							}
						}
					}
				}
			}
		}
	}

	return e
}

func calcMinM[T comparable](tBef Tree[T], tAft Tree[T], e [][][][][][]int) [][]int {
	lenBef := tBef.GetLength()
	lenAft := tAft.GetLength()

	minM := make([][]int, lenBef)
	for i := range minM {
		minM[i] = make([]int, lenAft)
	}

	minM[0][0] = 0
	for i := 1; i < lenBef; i++ {
		for j := 1; j < lenAft; j++ {
			minM[i][j] = 10000

			for _, s := range *tBef.GetPathToRoot(i) {
				for _, t := range *tAft.GetPathToRoot(j) {
					temp := minM[s][t] + e[s][tBef.GetNodes()[i].GetParent()][i-1][t][tAft.GetNodes()[j].GetParent()][j-1] - Cost(s, t, tBef, tAft)
					minM[i][j] = min(temp, minM[i][j])
				}
			}
			minM[i][j] = minM[i][j] + Cost(i, j, tBef, tAft)
		}
	}
	return minM
}

func calcD[T comparable](tBef Tree[T], tAft Tree[T], minM [][]int) [][]int {
	lenBef := tBef.GetLength()
	lenAft := tAft.GetLength()

	d := make([][]int, lenBef)
	for i := range minM {
		d[i] = make([]int, lenAft)
	}

	d[0][0] = 0

	for i := 1; i < lenBef; i++ {
		d[i][0] = d[i-1][0] + Cost(i, EmptyIndex, tBef, tAft)
	}

	for j := 1; j < lenAft; j++ {
		d[0][j] = d[0][j-1] + Cost(EmptyIndex, j, tBef, tAft)
	}

	for i := 1; i < lenBef; i++ {
		for j := 1; j < lenAft; j++ {
			d[i][j] = min(d[i][j-1]+Cost(EmptyIndex, j, tBef, tAft), d[i-1][j]+Cost(i, EmptyIndex, tBef, tAft), minM[i][j])
		}
	}

	return d
}

type Order struct {
	before  int
	after   int
	changed bool
}

func prepend[T any](s []T, v T) []T {
	return append([]T{v}, s...)
}

func calcOrder(d [][]int) []Order {
	order := make([]Order, 0)

	i := len(d) - 1
	j := len(d[i]) - 1

	for {
		if i == 0 && j == 0 {
			order = prepend(order, Order{before: i, after: j, changed: d[i][j] == 1})
			break
		}
		if i > 0 && j > 0 && d[i][j]-1 == d[i-1][j-1] {
			order = prepend(order, Order{i, j, true})
			i--
			j--
			continue
		}

		if j > 0 && d[i][j]-1 == d[i][j-1] {
			order = prepend(order, Order{EmptyIndex, j, true})
			j--
			continue
		}

		if i > 0 && d[i][j]-1 == d[i-1][j] {
			order = prepend(order, Order{i, EmptyIndex, true})
			i--
			continue
		}

		if i > 0 && j > 0 && d[i][j] == d[i-1][j-1] {
			order = prepend(order, Order{i, j, false})
			i--
			j--
			continue
		}

		break
	}

	return order
}

func CalcEditDistance[T comparable](treeBefore, treeAfter *model.Tree[T]) int {
	tBef := NewTree(treeBefore)
	tAft := NewTree(treeAfter)

	calcE[T](*tBef, *tAft)
	e := calcE[T](*tBef, *tAft)
	minM := calcMinM(*tBef, *tAft, e)
	d := calcD[T](*tBef, *tAft, minM)

	order := calcOrder(d)

	for _, o := range order {
		fmt.Println(o.before, o.after, o.changed)
	}
	return 0
}
