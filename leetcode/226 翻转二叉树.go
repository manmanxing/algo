package leetcode

/**
翻转一棵二叉树。

示例
输入
     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出
     4
   /   \
  7     2
 / \   / \
9   6 3   1

解题思路
先将根结点的左右节点翻转，再翻转根结点的左节点，再翻转根结点的右节点，这里使用递归，直到 root 节点为 nil 停止
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left,root.Right = root.Right,root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
