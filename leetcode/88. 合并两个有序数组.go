package leetcode

/**
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例一
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

示例二
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]

示例三
输入：nums1 = [0], m = 0, num2 = [1], n = 1
输出：[1]
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums2[n-1] >= nums1[m-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	//到这里有两种情况：1.m<=0,n>=0  2.m>0,n<=0
	for ; n-1 >= 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}