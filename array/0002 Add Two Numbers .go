package array

/**
2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。
示例
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

解题思路：
重点在 L1 和 L2 的迭代，利用 demo 对 L3 进行迭代更新
for循环终止条件是 L1 和 L2 都是nil，且没有余数
*/

type singleNode struct {
	Data int
	Next *singleNode
}

func addTwoNumbers(L1 *singleNode, L2 *singleNode) *singleNode {
	if L1 == nil {
		return L2
	}

	if L2 == nil {
		return L1
	}

	L3 := &singleNode{
		Data: 0,
		Next: nil,
	}
	demo := L3
	n1, n2, carry := 0, 0, 0
	for true {
		if L1 == nil {
			n1 = 0
		} else {
			n1 = L1.Data
			L1 = L1.Next
		}

		if L2 == nil {
			n2 = 0
		} else {
			n2 = L2.Data
			L2 = L2.Next
		}

		demo.Data = (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10

		if L1 == nil && L2 == nil && carry == 0 {
			return L3
		}

		demo.Next = &singleNode{
			Data: 0,
			Next: nil,
		}
		demo = demo.Next
	}

	return L3
}
