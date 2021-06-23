package structure

import "fmt"

/**
前序遍历的顺序是  根 -----> 左子树 -----> 右子树
中序遍历的顺序是 左子树 -----> 根 ------> 右子树
后序遍历的顺序是 左子树 -----> 右子树 -----> 根
*/

// TreeNode 二叉树结构
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}


func (t *TreeNode)Print()  {
	fmt.Print(t.Data, " ")
}

func (t *TreeNode)SetValue(data interface{})  {
	if t == nil {
		fmt.Println("tree node is nil")
		return
	}
	t.Data = data
}

// PreOrder 前序遍历
func (t *TreeNode) PreOrder() {
	if t == nil {
		return
	}
	t.Print()
	t.Left.PreOrder()
	t.Right.PreOrder()
}

// MiddleOrder 中序遍历
func (t *TreeNode) MiddleOrder() {
	if t == nil {
		return
	}
	t.Left.MiddleOrder()
	t.Print()
	t.Right.MiddleOrder()
}

// PostOrder 后序遍历
func (t *TreeNode) PostOrder() {
	if t == nil {
		return
	}
	t.Left.PostOrder()
	t.Right.PostOrder()
	t.Print()
}


//层次遍历(广度优先遍历)

