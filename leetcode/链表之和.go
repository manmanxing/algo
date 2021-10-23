package leetcode

/**
链表之和
限定语言：Python、C++、Java、Go、Javascript、Python 3
给定两个代表非负数的链表，数字在链表中是反向存储的（链表头结点处的数字是个位数，第二个结点上的数字是十位数...），求这个两个数的和，结果也用链表表示。
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出： 7 -> 0 -> 8
*/

func ListNodeSum(l1, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre.Next
	m := 0
	for l1 != nil || l2 != nil || m > 0 {
		sum := 0
		if m > 0 {
			sum ++
			m = 0
		}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum / 10 > 0 {
			m += 1
		}

		cur.Val = sum % 10
		cur = cur.Next
	}

	return pre.Next
}
