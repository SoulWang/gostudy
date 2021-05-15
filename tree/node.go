package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//显示定义和命名方法接收者
func (node Node) Print() {
	fmt.Print(node.Value)
}

//使用指针作为方法接收者，只有指针才可以改变结构内容，nil指针也可以调用方法
//要改变内容必须使用指针接收者
//结构过大也考虑使用指针接收者
//一致性：如有指针接收者，最好都是指针接收者
//值接收者：为go语言特有的
//值/指针接收者均可接受值/指针
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node.Ignord")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

/*使用了自定义工厂函数，注意返回了局部变量的地址*/
func CreateNode(value int) *Node {
	return &Node{Value: value}

}
