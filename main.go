package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func main() {
	tree := model.NewTree("root")

	c1 := model.NewNode("child1")
	c2 := model.NewNode("child2")

	tree.Root.AddChild(c1)
	tree.Root.AddChild(c2)

	c11 := model.NewNode("child1.1")
	c12 := model.NewNode("child1.2")
	c21 := model.NewNode("child2.1")

	c1.AddChild(c11)
	c1.AddChild(c12)
	c2.AddChild(c21)

	fmt.Println(tree)
}
