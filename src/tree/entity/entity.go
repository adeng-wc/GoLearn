package main

import (
	"fmt"
	"GoLearn/src/tree"
)

// 组合
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()
	root.SetValue(2)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
