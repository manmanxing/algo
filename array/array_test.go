package array

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
