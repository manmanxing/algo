package leetcode


/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例一
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

解题思路：
双指针left与right往中间靠拢，高度每到 high 限制时，就要计算该层有多少个单位，需要将水填充的高度也看作柱子
一直循环到双指针 left > right，这样就计算了水填充的单位+柱子所占单位的总和
 */



func trap(height []int) int {
	//high 代表是哪一层，初始化为第一层
	//left，right 分别代表左边第一个柱子，右边第一个柱子
	left,right,high:=0,len(height)-1,1
	//temp 填满水后所有柱子(将水也看作柱子的一部分)的单位个数
	//sum 所有柱子(不包含水)的单位个数
	//temp - sum 就是代表多少个单位的水量
	temp,sum := 0,0
	for left <= right{
		for left<=right && height[left]<high{
			//第一次循环，如果左边的柱子高度<1，那么肯定要往右移
			left ++
		}
		for right >= left && height[right]<high{
			//第一次循环，如果右边的柱子高度<1，那么肯定要往左移
			right --
		}
		//到这里，表示在该层 left，right下表所在的高度都已经 >= high
		//计算在该层能有多少个柱子(将水也看作柱子的一部分)
		temp += right-left + 1
		//该层遍历完毕，层数加一
		high ++
	}

	for _,i := range height{
		sum += i
	}

	return temp - sum
}