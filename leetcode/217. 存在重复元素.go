package leetcode

import "sort"

/**
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

示例 1:
输入: [1,2,3,1]
输出: true

示例 2:
输入: [1,2,3,4]
输出: false

示例 3:
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/

//利用map
func containsDuplicate1(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	//利用 struct{}，降低内存消耗
	var demo = make(map[int]struct{}, 0)
	for i := range nums {
		if _, ok := demo[nums[i]]; ok {
			return true
		}
		demo[nums[i]] = struct{}{}
	}
	return false
}

//排序
func containsDuplicate2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sort.Ints(nums)
	demo := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == demo {
			return true
		}else {
			demo = nums[i]
		}
	}
	return false
}