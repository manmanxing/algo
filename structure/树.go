package structure

import "fmt"

/**
前序遍历的顺序是 先访问当前结点，再前序遍历左子树，最后再前序遍历右子树，即根—左—右。
中序遍历的顺序是 先中序遍历左子树，然后再访问当前结点，最后再中序遍历右子树，即左—根—右。
后序遍历的顺序是 先后序遍历左子树，然后再后序遍历右子树，最后再访问当前结点，即左—右—根。
层次遍历的顺序是 从第一层开始，从左到右依此遍历每层，直到结束。
*/

// TreeNode 二叉树结构
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}


func (t *TreeNode)Print()  {
	fmt.Print(t.Data, " ")
}

func (t *TreeNode) SetValue(data int) {
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

//层次遍历(广度优先遍历 BFS)
func (t *TreeNode) BreadthFirstSearch() {
	if t == nil {
		return
	}

	//用于装载节点的数据
	result := make([]int, 0)
	//用于装载root，左右节点
	//顺序为父节点a，a的左节点，a的右节点；父节点b，b的左节点，b的右节点；
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, t)

	for len(nodes) > 0 {
		//获取nodes 里的第一个节点
		headNode := nodes[0]
		//保存nodes 里的第一个节点的数据
		result = append(result, headNode.Data)
		//删除nodes的第一个节点
		nodes = nodes[1:]

		//记录每层的节点
		//遍历顺序是 从第一层开始，从左到右依此遍历每层，直到结束。
		//因此先获取该父节点下左边的节点，再获取右边的节点
		if headNode.Left != nil {
			nodes = append(nodes, headNode.Left)
		}
		if headNode.Right != nil {
			nodes = append(nodes, headNode.Right)
		}
	}

	//装载数据完毕，打印
	for i := range result {
		fmt.Print(result[i]," ")
	}
}

//获取树的层数（使用递归）
//对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
func (t *TreeNode) LayersByRecursion() int {
	if t == nil {
		return 0
	}
	leftLayers := t.Left.LayersByRecursion()
	rightLayers := t.Right.LayersByRecursion()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}


//获取树的层数（不使用递归）
func (t *TreeNode) Layers() int  {
	var layer int

	if t == nil {
		return  layer
	}

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, t)

	for len(nodes) > 0 {
		layer ++
		//获取该层节点个数
		size := len(nodes)
		for size > 0 {
			//获取第一个节点，然后在nodes删除该节点
			headNode := nodes[0]
			nodes = nodes[1:]
			//然后 nodes 装载下一层的节点
			if headNode.Left != nil {
				nodes = append(nodes, headNode.Left)
			}
			if headNode.Right != nil {
				nodes = append(nodes,headNode.Right)
			}
			size --
		}
	}

	return layer
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{Data: v}
}