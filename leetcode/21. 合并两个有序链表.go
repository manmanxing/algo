package leetcode

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例一
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例二
输入：l1 = [], l2 = []
输出：[]

示例三
输入：l1 = [], l2 = [0]
输出：[0]

*/


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	//使用递归，不断去找两个链表中比较小的元素，然后result接上那个元素
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
