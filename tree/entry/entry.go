package main

import (
	"fmt"
	"learning/tree"
)

type TreeNode struct {
	node *tree.Node
}

func (treeNode *TreeNode) postOrder() {
	if treeNode == nil || treeNode.node == nil {
		return
	}

	left := TreeNode{treeNode.node.Left}
	right := TreeNode{treeNode.node.Right}
	right.postOrder()
	left.postOrder()
	treeNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{
		Value: 3,
	}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(6)
	root.Print()
	root.Right.Left.SetValue(55)
	root.Right.Left.Print()
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(root)
	fmt.Println(nodes)
}
