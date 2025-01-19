package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
)

func main() {
	tree := model.NewTree[string](model.NewLabel("root"))

	c1 := model.NewNodeFromValue("child1")
	c2 := model.NewNodeFromValue("child2")

	tree.Root.AddChild(c1)
	tree.Root.AddChild(c2)

	c11 := model.NewNodeFromValue("child1.1")
	c12 := model.NewNodeFromValue("child1.2")
	c21 := model.NewNodeFromValue("child2.1")

	c1.AddChild(c11)
	c1.AddChild(c12)
	c2.AddChild(c21)

	fmt.Println(tree)
}
