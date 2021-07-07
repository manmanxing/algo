package leetcode

/**
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

示例一
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例二
输入：head = [2,1], x = 2
输出：[1,2]
*/

//要对链表做修改，需要对头节点做一个复制，必须对头节点的下一个节点做修改，而不是从头节点开始修改
//比如small为头节点的链表，对头节点做一个复制：s := small
//利用对 s.next 的修改来间接修改 small，这样 small 的next节点也会跟着修改
func partition(head *ListNode, x int) *ListNode {
	small, big := new(ListNode), new(ListNode)
	s, b := small, big
	for head != nil {
		if head.Val < x {
			s.Next = head //复用listNode，而不是创建新的listNode
			//1-4-3-2-5-2
			//2-5-2
			//2(最后一次 s的链表情况)
			s = s.Next
		} else {
			//4-3-2-5-2
			//3-2-5-2
			//5-2(最后一次 b的链表情况)
			b.Next = head
			b = b.Next
		}
		head = head.Next
	}
	//此时，s是head里最后一个小于3的数以及后面的节点，b 是head里最后一个大于等于3的数以及后面的节点
	//注意：s和b一定有交集：要么 s包含于 b 中(s:2,b:5-2)，要么 b 包含于 a 中(s:2-5,b:5)
	//设置 small 的最后一个节点 s 的 next 是 big的 next节点
	//设置 big 的最后一个节点 b 为 nil
	b.Next = nil
	s.Next = big.Next
	return small.Next
}