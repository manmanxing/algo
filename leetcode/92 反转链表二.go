package leetcode


/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
这个区别在于：只需要翻转位置 left 到位置 right 的链表节点，而不是所有的节点。

示例1
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例2
输入：head = [5], left = 1, right = 1
输出：[5]

解题思路：
虚拟节点和头插法
以[1,2,3,4,5]为例，首先将将2后面的3移到1的后面成为 [1,3,2,4,5]
再将2后面的4移到1的后面成为[1,4,3,2,5]
因此需要两个指针m,n，一个表示1，一个表示2
 */


func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := new(ListNode)
	pre.Next = head

	m,n := pre,pre.Next

	//根据 left 确定 m，n的具体位置
	for left-1 > 0 {
		m = m.Next
		n = n.Next
	}

	//头插法
	sub := right - left
	for sub > 0 {
		m.Next,n.Next,n.Next.Next = n.Next,n.Next.Next,m.Next
		sub --
	}

	return pre.Next
}