package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node // * 指针类型
}

// Node 接受者    == func  print(Node Node)
// 值接收者
func (node Node) Print() {
	fmt.Println(node.Value)
}

// *Node 指针接受者
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting Value is nil")
		return
	}
	node.Value = value
}

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} // 创建局部变量地址也能给外部使用。  结果创建在堆上还是栈上，不需要知道
}

// 遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
