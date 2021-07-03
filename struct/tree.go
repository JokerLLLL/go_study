package main

import (
	"fmt"
	"github.com/jokerl/go_study/struct/tree"
)

func main()  {
	node := tree.Node{Value:4}
	node.Left = &tree.Node{Value:2}
	node.Right = &tree.Node{Value:3,Left:&tree.Node{1,nil,nil}}
	node.Right.Right = new(tree.Node) // new 返回Node的指针

	/**
	       node(4)
	      /      \
	  node(5)   node(3)
	            /    \
	        node(1)  node(0)
	*/
	node.Left.SetValue(5)
	node.Left.Print()
	poitNode := &node
	poitNode.Print()
	poitNode = nil
	node.Left.Left.Show() // nil 指针不报错

	println()
	node.Traverse()
	println()
	node.TraversBefore()
	println()
	myNode := tree.MyNode{&node}
	myNode.TraversAfter()
	println()

	count := 0
	node.TraversFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println("tree count:", count)

	max := 0
	c := node.TraversChanel()
	for node := range c{
		if node.Value > max {
			max = node.Value
		}
	}
	fmt.Printf("max: %d", max)

}