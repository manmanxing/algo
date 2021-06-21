package leetcode

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例一
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例二
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

示例三
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000

示例四
输入：nums1 = [], nums2 = [1]
输出：1.00000

示例五
输入：nums1 = [2], nums2 = []
输出：2.00000
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var middle float64
	nums1 = append(nums1, nums2...)
	middle = findMedianSortedArraysnext(nums1)
	return middle
}

func findMedianSortedArraysnext(list []int) float64 {
	if len(list) <= 0 {
		return 0
	}

	if len(list) == 1 {
		return float64(list[0])
	}

	start, end := list[0], list[len(list)-1]
	for {

	}
}
