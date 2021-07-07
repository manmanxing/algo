package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(nums, target))
}

func TestAddTwoNumbers(t *testing.T) {
	L1 := &singleNode{
		Data: 2,
		Next: nil,
	}
	L1.Next = &singleNode{
		Data: 4,
		Next: nil,
	}

	L1.Next.Next = &singleNode{
		Data: 3,
		Next: nil,
	}

	L2 := &singleNode{
		Data: 5,
		Next: nil,
	}
	L2.Next = &singleNode{
		Data: 6,
		Next: nil,
	}

	L2.Next.Next = &singleNode{
		Data: 4,
		Next: nil,
	}

	L3 := addTwoNumbers(L1, L2)
	var str string
	for L3 != nil {
		str += strconv.Itoa(L3.Data)
		if L3.Next != nil {
			str += "->"
		}
		L3 = L3.Next
	}
	fmt.Println("L3:", str)
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	a := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring1(a))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	a := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(a))
}

func TestFindMedianSortedArrays(t *testing.T) {
	num1, num2 := []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

func TestGetMaxAndSecondNum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(getMaxAndSecondNum(a))
}

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	cur := head
	for i := 1; i < 10; i++ {
		demo := &ListNode{
			Val:  i+1,
			Next: nil,
		}
		cur.Next = demo
		cur = cur.Next
	}
	print(head)
	head = reverseKGroup(head, 4)
	print(head)
}

func print(head *ListNode)  {
	if head == nil {
		return
	}
	cur := head
	result := ""
	for cur != nil {
		result += strconv.Itoa(cur.Val)
		if cur.Next != nil {
			result += "=>"
		}
		cur = cur.Next
	}
	fmt.Println(result)
}

func TestAddStrings(t *testing.T)  {
	fmt.Println(addStrings("99898","3489"))
}

func TestPartition(t *testing.T)  {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	a := []int{1,4,3,2,2,5}
	cur := head//对head 头节点进行复制
	for i := range a {
		demo := &ListNode{
			Val:  a[i],
			Next: nil,
		}
		//对头节点的 next 进行修改
		cur.Next = demo
		cur = cur.Next
	}
	//最终 head.next 是我们需要的链表
	print(head.Next)
	head = partition(head.Next,3)
	print(head)
}