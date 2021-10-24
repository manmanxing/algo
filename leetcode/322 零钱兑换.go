package leetcode

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例1
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例2
输入：coins = [2], amount = 3
输出：-1

示例3
输入：coins = [1], amount = 0
输出：0

示例4
输入：coins = [1], amount = 1
输出：1

示例5
输入：coins = [1], amount = 2
输出：2

使用动态规划
比如 coins = [1, 2, 5], amount = 11
题目求的值为 f(11)，第一次选择硬币时我们有三种选择。
假设我们取面额为 1 的硬币，那么接下来需要凑齐的总金额变为 11 - 1 = 10，即 f(11) = f(10) + 1，这里的 +1 就是我们取出的面额为 1 的硬币。
同理，如果取面额为 2 或面额为 5 的硬币可以得到：
    f(11) = f(9) + 1
    f(11) = f(6) + 1
所以：
f(11) = min(f(10), f(9), f(6)) + 1

假设 f(n) 代表要凑齐金额为 n 所要用的最少硬币数量，那么有：
f(n) = min(f(n - c1), f(n - c2), ... f(n - cn)) + 1
其中 c1 ~ cn 为硬币的所有面额。
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				//求最小值
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if amount < dp[amount] {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
