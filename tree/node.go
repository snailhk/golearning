package tree

import "fmt"

// 要用指针接受者
// 结构过大也考虑使用指针接收者
// 一致性: 如有指针接收者，最好都是指针接收者

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
