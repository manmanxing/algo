package leetcode

import "sort"

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？
请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例一
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例二
输入：nums = []
输出：[]

示例三
输入：nums = [0]
输出：[]

思路：
外层循环：指针 i 遍历数组。
内层循环：用双指针，去寻找满足三数之和 == 0 的元素

先排序的意义
便于跳过重复元素，如果当前元素和前一个元素相同，跳过。

双指针的移动时，避免出现重复解
找到一个解后，左右指针同时向内收缩，为了避免指向重复的元素，需要：

左指针在保证left < right的前提下，一直右移，直到指向不重复的元素
右指针在保证left < right的前提下，一直左移，直到指向不重复的元素
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	//i < len(nums)-2：当元素个数<=3时，不需要执行
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		//如果最小的都大于0,break
		if n1 > 0 {
			break
		}
		//如果和前一个相同，跳过
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		//转换为了两数之和
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				//=0,加入结果集
				res = append(res, []int{n1, n2, n3})
				//去重且移位
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 < 0 {
				//<0,left右移
				left++
			} else {
				//>0,right左移
				right--
			}
		}
	}
	return res
}
