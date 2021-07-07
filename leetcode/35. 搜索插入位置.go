package leetcode

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0

*/

//一个一个的判断
func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	//说明 target 是最大值
	return len(nums)
}

//下标记录
func searchInsert2(nums []int, target int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			index++
		} else {
			return index
		}
	}
	return index
}

//二分法, nums 是有序的
//需要注意边界问题
func searchInsert3(nums []int, target int) int {
	left,right := 0,len(nums)

	for left < right {
		middle := (left+right)/2
		if nums[middle] >= target {
			right = middle
		}else {
			left = middle+1
		}
	}
	return left
}