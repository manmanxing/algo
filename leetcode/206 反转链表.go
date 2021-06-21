package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre,cur *ListNode= nil,head

	for cur != nil {
		pre,cur,cur.Next = cur,cur.Next,pre
	}

	return pre
}
