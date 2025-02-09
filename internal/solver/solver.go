package solver

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
	"math"
)

func calcE[T any](tBef Tree[T], tAft Tree[T]) [][][][][][]int {
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
			for _, u := range *tBef.GetPathToRoot(i) {
				for _, s := range *tBef.GetPathToRoot(u) {
					for _, v := range *tAft.GetPathToRoot(j) {
						for _, t := range *tAft.GetPathToRoot(v) {
							if (s == u && u == i) && (t == v && v == j) {
								e[s][u][i][t][v][j] = Cost(i, j)
							} else if (s == u && u == i) || (t < v && v == j) {
								e[s][u][i][t][v][j] = e[s][u][i][t][tAft.GetNodes()[j].GetParent()][j-1] + Cost(EmptyIndex, j)
							} else if (s < u && u == i) || (t == v && v == j) {
								e[s][u][i][t][v][j] = e[s][tBef.GetNodes()[i].GetParent()][i-1][t][v][j] + Cost(i, EmptyIndex)
							} else {
								x := tBef.GetChildOnPath(u, i)
								y := tAft.GetChildOnPath(v, j)
								e[s][u][i][t][v][j] = min(e[s][x][i][t][v][j], e[s][u][i][t][y][j], e[s][u][x-1][t][v][y-1], e[x][x][i][y][y][j])
							}
						}
					}
				}
			}
		}
	}

	return e
}

func calcMinM[T any](tBef Tree[T], tAft Tree[T], e [][][][][][]int) [][]int {
	lenBef := tBef.GetLength()
	lenAft := tAft.GetLength()

	minM := make([][]int, lenBef)
	for i := range minM {
		minM[i] = make([]int, lenAft)
	}

	minM[0][0] = 0
	for i := 1; i < lenBef; i++ {
		for j := 1; j < lenAft; j++ {
			minM[i][j] = math.MaxInt

			for _, s := range *tBef.GetPathToRoot(i) {
				for _, t := range *tAft.GetPathToRoot(j) {
					temp := minM[s][t] + e[s][tBef.GetNodes()[i].GetParent()][i-1][t][tAft.GetNodes()[j].GetParent()][j-1] - Cost(s, t)
					minM[i][j] = min(temp, minM[i][j])
				}
			}
			minM[i][j] = minM[i][j] + Cost(i, j)
		}
	}
	return minM
}

func calcD[T any](tBef Tree[T], tAft Tree[T], minM [][]int) [][]int {
	lenBef := tBef.GetLength()
	lenAft := tAft.GetLength()

	d := make([][]int, lenBef)
	for i := range minM {
		d[i] = make([]int, lenAft)
	}

	d[0][0] = 0

	for i := 1; i < lenBef; i++ {
		d[i][0] = d[i-1][0] + Cost(i, EmptyIndex)
	}

	for j := 1; j < lenAft; j++ {
		d[0][j] = d[0][j-1] + Cost(EmptyIndex, j)
	}

	for i := 1; i < lenBef; i++ {
		for j := 1; j < lenAft; j++ {
			d[i][j] = min(d[i][j-1]+Cost(EmptyIndex, j), d[i-1][j]+Cost(i, EmptyIndex), minM[i][j])
		}
	}

	return d
}

func CalcEditDistance[T any](treeBefore, treeAfter *model.Tree[T]) int {
	tBef := NewTree(treeBefore)
	tAft := NewTree(treeAfter)

	e := calcE[T](*tBef, *tAft)

	fmt.Println(e)

	minM := calcMinM(*tBef, *tAft, e)

	fmt.Println(minM)

	d := calcD[T](*tBef, *tAft, minM)

	fmt.Println(d)

	return 0
}
