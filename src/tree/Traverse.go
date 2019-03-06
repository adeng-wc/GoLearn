package tree

import "fmt"

// 中序 遍历
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}
