package leetcode

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
最近公共祖先的定义为：对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

示例一
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

示例二
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

示例三
输入：root = [1,2], p = 1, q = 2
输出：1

解题思路：
首先明白最近公共祖先必须得让p、q分布它的异侧。
函数代表root是否等于p或q，递归调用root.left返回结果表示左子树是否包含p或q，
若递归调用的root.left为空表示p和q都在右侧，若root.right为空表示p和q都在左侧，这两种情况都需要继续递归，
若两者都不为空则符合定义，代表root就是最近的公共祖先。

 */


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == p.Data || root.Data == q.Data {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}