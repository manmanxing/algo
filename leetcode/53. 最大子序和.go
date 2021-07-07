package leetcode

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例一
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例二
输入：nums = [1]
输出：1

解题思路：
如果前边累加后还不如自己本身大，说明前面的值为负数，那就把前边的都扔掉，从此自己本身重新开始累加。
*/

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	//从第二位开始，跟前面一位进行求和
	for i := 1; i < len(nums); i++ {
		//和比自己大，那就累加
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		//记录每次累计中的最大值
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}