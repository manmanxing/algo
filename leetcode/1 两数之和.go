package leetcode

/**
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
示例：
给定 nums = [2, 7, 11, 15], target = 9,
因为 nums[0] + nums[1] = 2 + 7 = 9,
输出 [0, 1]

解题思路：
因为 target - nums[0] = nums[1]
遍历 nums, 将 target -  nums[i] 的差保存到 map的key，value寸对应的索引
保存之前需判断该差是否已经在map存在，若存在，说明找到了，并输出value即可
*/

func TwoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return nil
	}
	demoMap := make(map[int]int, 0)
	index := make([]int, 0)
	for k, v := range nums {
		d := target - v
		if i, ok := demoMap[d]; ok {
			index = append(index, i, k)
			return index
		}
		demoMap[v] = k
	}

	return index
}
