package main

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left, right *TreeNode // * 指针类型
}

// node 接受者    == func  print(node TreeNode)
// 值接收者
func (node TreeNode) print() {
	fmt.Println(node.value)
}

// *TreeNode 指针接受者
func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value is nil")
		return
	}
	node.value = value
}

// 工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{value: value} // 创建局部变量地址也能给外部使用。  结果创建在堆上还是栈上，不需要知道
}

// 遍历
func (node *TreeNode) traverse()  {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	root := TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)
	root.left.right = createNode(2)

	root.traverse()



	nodes := []TreeNode{
		{value: 3},
		{},
		{5, nil, &root},
	}

	fmt.Println(nodes)
	root.setValue(2)
	root.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

}
