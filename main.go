package main

import (
	"github.com/Hietan/TreeEditDistance/internal/parser"
	"github.com/Hietan/TreeEditDistance/internal/solver"
)

func main() {
	tree1 := parser.LoadTreeFromFile[string]("./tree1.json")
	tree2 := parser.LoadTreeFromFile[string]("./tree2.json")

	solver.CalcEditDistance(tree1, tree2)
}
