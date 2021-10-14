package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//这里注意：不能写成 pre,cur *ListNode := nil,head。因为 nil 不能单独复制给未指定类型的变量
	var pre,cur *ListNode= nil,head

	for cur != nil {
		pre,cur,cur.Next = cur,cur.Next,pre
	}

	return pre
}
