package leetcode

import "sort"

/**
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的 最短 子数组，并输出它的长度。

示例1
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

示例2
输入：nums = [1,2,3,4]
输出：0

示例3
输入：nums = [1]
输出：0

解题思路
我们将给定的数组 nums 表示为三段子数组拼接的形式，分别记作 numsA，numsB，numsC。当我们对 numsB 进行排序，整个数组将变为有序。
换而言之，当我们对整个序列进行排序，numsA 和 numsC 都不会改变。

本题要求我们找到最短的 numsB，即找到最大的 numsA 和 numsC 的长度之和。
因此我们将原数组 nums 排序与原数组进行比较，取最长的相同的前缀为 numsA，取最长的相同的后缀为 numsC，这样我们就可以取到最短的 numsB。

具体地，我们创建数组 nums 的拷贝，记作数组 numsSorted，并对该数组进行排序，然后我们从左向右找到第一个两数组不同的位置，即为 numsB 的左边界。
同理也可以找到 numsB 右边界。最后我们输出 numsB 的长度即可。

特别地，当原数组有序时，numsB 的长度为 000，我们可以直接返回结果。
*/

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	numsSorted := append([]int{}, nums...)
	sort.Ints(numsSorted)

	//从左右两边开始找
	left, right := 0, len(nums)-1

	for nums[left] == numsSorted[left] {
		left++
	}

	for nums[right] == numsSorted[right] {
		right--
	}

	return right - left + 1
}
