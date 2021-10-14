package leetcode

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例一
输入：head = [1,3,2]
输出：[2,3,1]

解题思路一：
因为需要从尾到头打印，而正常for 循环是从头到尾，只需要利用 defer 先进后出特点就可以。

解题思路二：
append 将 val 放在前头
 */


func reversePrint1(head *ListNode) (result []int) {
	//var result []int
	for head != nil {
		curVal := head.Val
		defer func() {
			result = append(result, curVal)
		}()
		head = head.Next
	}
	return
}

func reversePrint2(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append([]int{head.Val},result...)
		head = head.Next
	}
	return result
}

