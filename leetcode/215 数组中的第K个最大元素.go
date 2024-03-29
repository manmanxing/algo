package leetcode

import "sort"

/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4  //注意，这里没有去重
输出: 4
 */


func findKthLargest(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}

	sort.Ints(nums)

	return nums[len(nums)-k]
}