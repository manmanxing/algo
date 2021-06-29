package leetcode

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例一
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回其层序遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
 */

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	//nodes 装载的是对应层的节点
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		demo := make([]int,0)//用于装载该层的数据
		size := len(nodes)
		//遍历该层的所有节点
		for size >0 {
			headNode := nodes[0]
			demo = append(demo, headNode.Data)
			nodes = nodes[1:]

			if headNode.Left != nil {
				nodes = append(nodes, headNode.Left)
			}

			if headNode.Right != nil {
				nodes = append(nodes, headNode.Right)
			}

			size --
		}

		result = append(result, demo)
	}

	return result
}
