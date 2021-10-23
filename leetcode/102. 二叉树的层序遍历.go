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

解题思路：
遍历每一层的节点，保存每一层节点的值
使用数组A将需要遍历的节点保存起来，比如刚开始是从 root 节点开始的，在遍历过程中数组A会将 root 的左右子节点也保存。
只要删除数组A的第一个节点，就可以顺移到下一个节点继续遍历
*/

func levelOrder1(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	//用于记录该层的所有节点
	curNode := make([]*TreeNode, 0)
	//首先从根结点开始
	curNode = append(curNode, root)

	for len(curNode) > 0 {
		//该层节点元素组成的数组
		curArray := make([]int, 0)
		for i := 0; i < len(curNode); i++ {
			curArray = append(curArray, curNode[i].Data)
			if curNode[i].Left != nil {
				curNode = append(curNode, curNode[i].Left)
			}

			if curNode[i].Right != nil {
				curNode = append(curNode, curNode[i].Right)
			}
			//移除第一个curNode
			curNode = curNode[1:]
		}
		//将该层node 的节点数组保存起来
		result = append(result, curArray)
	}

	return result
}

func levelOrder2(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	//nodes 装载的是对应层的节点
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		demo := make([]int, 0) //用于装载该层的数据
		size := len(nodes)
		//遍历该层的所有节点
		for size > 0 {
			headNode := nodes[0]
			demo = append(demo, headNode.Data)
			nodes = nodes[1:]

			if headNode.Left != nil {
				nodes = append(nodes, headNode.Left)
			}

			if headNode.Right != nil {
				nodes = append(nodes, headNode.Right)
			}

			size--
		}

		result = append(result, demo)
	}

	return result
}
