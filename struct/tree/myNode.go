package tree

// 组合新增方法
type MyNode struct {
	Node *Node
}

func (myNode *MyNode) TraversAfter()  {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := MyNode{myNode.Node.Left}
	right := MyNode{myNode.Node.Right}
	left.TraversAfter()
	right.TraversAfter()
	myNode.Node.Print()
}

func (myNode MyNode) DoSomeThing() {
	//TODO..
}
