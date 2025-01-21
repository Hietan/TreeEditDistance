package main

import (
	"fmt"
	"github.com/Hietan/TreeEditDistance/internal/model"
	"github.com/Hietan/TreeEditDistance/internal/solver"
)

func main() {
	tree1 := model.NewTree("root")

	c1 := model.NewNode("child1")
	c2 := model.NewNode("child2")

	tree1.GetRoot.AddChild(c1)
	tree1.GetRoot.AddChild(c2)

	c11 := model.NewNode("child1.1")
	c12 := model.NewNode("child1.2")
	c21 := model.NewNode("child2.1")

	c1.AddChild(c11)
	c1.AddChild(c12)
	c2.AddChild(c21)

	fmt.Println("Tree1\n", tree1)

	tree2 := model.NewTree("root")

	d1 := model.NewNode("child1")
	d2 := model.NewNode("child2")

	tree2.GetRoot.AddChild(d1)
	tree2.GetRoot.AddChild(d2)

	d11 := model.NewNode("child1.1")
	d12 := model.NewNode("child1.2")
	d21 := model.NewNode("child2.1")

	d1.AddChild(d11)
	d1.AddChild(d12)
	d2.AddChild(d21)

	fmt.Println("Tree2\n", tree2)

	distance := solver.CalcEditDistance(tree1, tree2)
	fmt.Println("Edit distance:", distance)
}
