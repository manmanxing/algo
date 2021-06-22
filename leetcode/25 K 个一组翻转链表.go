package leetcode


/**
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

示例一
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例二
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

示例三
输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]

示例四
输入：head = [1], k = 1
输出：[1]
 */


func reverseKGroup(head *ListNode, k int) *ListNode {
	//设置一个起点
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}

	pre := dummy
	cur := dummy.Next

	for {
		n := k
		//先找出下一批次的头节点
		nextCur := cur
		for n> 0 && nextCur != nil {
			n --
			nextCur = nextCur.Next
		}
		//到这里，有两种情况
		//1. n>0 ,nextCur = nil,表示到了最后了
		//2. n <=0, nextCur !=nil，表示还没到最后
		if n > 0 {
			break
		}else {
			//重新让 n = k,是为了方便这k个节点反转
			n = k
		}

		//保留下一批次的 pre 节点
		nextPre := cur

		//k个节点反转
		for n > 0 {
			nextCur,cur,cur.Next = cur,cur.Next,nextCur
			n--
		}
		//反转完毕后，重新组装当前链表
		pre.Next = nextCur
		pre = nextPre
		cur = pre.Next
	}

	return dummy.Next
}