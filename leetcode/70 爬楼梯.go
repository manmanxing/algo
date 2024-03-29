package leetcode

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例1
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例2
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

解题思路1
第n个台阶只能从第n-1或者n-2个上来。
到第n-1个台阶的走法 + 第n-2个台阶的走法 = 到第n个台阶的走法
已经知道了第1个和第2个台阶的走法，一路加上去。
dp[n] = dp[n-1] + dp[n-2]

1=1
2=2+1
3=3+2+1
4=4+3+2+1
*/

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	//a 表示第n-2个台阶的走法,b表示第n-1个台阶的走法,sum表示第 N 阶台阶需要的走法
	a, b, sum := 1, 2, 0
	for i := 3; i <= n; i++ {
		//累加结果就是第n个台阶的走法
		sum = a + b
		//向下迭代
		a = b   //下次迭代的第n-2个台阶的走法等于上次迭代n-1个台阶的走法
		b = sum //下次迭代的第n-1个台阶的走法等于上次迭代的第n个台阶走法
	}

	return sum
}
