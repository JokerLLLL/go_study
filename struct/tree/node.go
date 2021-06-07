package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

//工厂函数 === construct
func CreateNode(value int) *Node  {
	return &Node{Value:value}
}

func (node *Node) Show() {
	fmt.Println("show...")
}

func (node *Node) Print() {
	if node == nil {
		fmt.Println("this is nil pointer")
		return
	}
	print(node.Value)
}

func (node *Node) SetValue(value int) *Node {
	node.Value = value
	return node
}

func (node *Node) Traverse()  {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) TraversBefore()  {
	if node == nil {
		return
	}
	node.Print()
	node.Left.TraversBefore()
	node.Right.TraversBefore()
}

// 中序遍历处理闭包
func (node *Node) TraversFunc(f func(node *Node))  {
	if node == nil {
		return
	}
	node.Left.TraversFunc(f)
	f(node)
	node.Right.TraversFunc(f)
}



