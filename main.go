package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/parser"
	"github.com/Hietan/TreeEditDistance/internal/solver"
)

func main() {
	tree1 := parser.LoadTreeFromFile[string]("./tree1.json")
	tree2 := parser.LoadTreeFromFile[string]("./tree2.json")

	fmt.Println(tree1)
	fmt.Println(tree2)

	solver.CalcEditDistance(tree1, tree2)
}
