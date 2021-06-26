package leetcode

/**
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回锯齿形层序遍历如下：
[
  [3],
  [20,9],
  [15,7]
]

方法一：使用广度优先遍历 BFS
 */


type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int//记录每层的数据

	if root == nil {
		return ret
	}

	//记录每层的节点
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	//记录是否需要反转遍历方向
	flag := false

	for len(nodes) > 0 {
		//获取该层节点个数
		size := len(nodes)
		//记录该层的节点数据
		curListNodes := make([]int,0)
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
			curListNodes = append(curListNodes, headNode.Data)
			size --
		}
		//根据 flag 判断是否需要反转
		if flag {
			reverse(curListNodes)
		}
		ret = append(ret, curListNodes)
		flag = !flag
	}

	return ret
}

//数组原地反转
func reverse(a []int){
	right,left := len(a)-1,0
	for left < right {
		temp := a[right]
		a[right] = a[left]
		a[left] = temp
		right--
		left++
	}
	return
}