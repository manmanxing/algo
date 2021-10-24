package leetcode

/**
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例1
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例2
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

解题思路
双指针解法
用两个指针p1，p2分别指向第一个节点与第二个节点
p1的下一个节点是p2的下一个节点

1->2->3->4->5->NULL
single -> 1
double -> 2
------------------
1.Next -> 2.Next
1 -> 3
2.Next -> 3.Next
2 -> 4
------------------
3.Next -> 4.Next
3 -> 5
4.Next -> 5.Next
4 -> 6
*/

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//head 代表奇数链表，ouHead 代表偶数链表
	//用两个指针p1，p2分别指向第一个奇数节点与第一个偶数节点
	p1, p2, ouHead := head, head.Next, head.Next

	for p1.Next != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next

		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = ouHead
	return head
}
