package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/parser"
)

func main() {
	tree := parser.LoadTreeFromFile[string]("./tree1.json")
	fmt.Println(tree)
}
