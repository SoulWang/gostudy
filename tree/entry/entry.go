package main

import (
	"fmt"
	"gostudy/tree"
)

/**
不论地址还是结构本身，一律使用.来访问成员
go仅支持封装，不支持继承和多态
go语言没有class，只有struct
*/
/**
结构创建在堆上还是栈上
在java中几乎所有的东西都是分配在堆上
在go中，我们不需要知道treeNode到底分配在哪里
由go语言的编译器和运行环境决定的
*/

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
