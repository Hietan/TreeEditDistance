package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/parser"
	"github.com/Hietan/TreeEditDistance/internal/solver"
)

func main() {
	tree := parser.LoadTreeFromFile[string]("./tree1.json")
	fmt.Println(tree)

	solver.CalcEditDistance(tree, tree)
}
