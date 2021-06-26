package structure

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {
	root := TreeNode{Data: 3}
	root.Left = &TreeNode{}
	root.Left.SetValue(0)
	root.Left.Right = CreateNode(2)
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = CreateNode(4)

	fmt.Print("\n前序遍历: ")
	root.PreOrder()
	fmt.Print("\n中序遍历: ")
	root.MiddleOrder()
	fmt.Print("\n后序遍历: ")
	root.PostOrder()
	fmt.Print("\n层次遍历: ")
	root.BreadthFirstSearch()
	fmt.Print("\n层数: ", root.LayersByRecursion())
	fmt.Println("\n层数: ", root.Layers())
}