package leetcode

import "sort"

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例一
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例二
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

示例三
输入：intervals = [[1,4],[1,4]]
输出：[[1,4]]

示例四
输入：intervals =  [[1,4],[0,4]]
输出：[[0,4]]

示例五
输入：intervals =  [[1,4],[2,3]]
输出：[[1,4]]


提示：
1.任意单个区间i 的长度都为2，且 0 <= starti <= endi
*/


func mergeRepeatArray(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return intervals
	}

	//先排序,左边界还是右边界都可以
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := 0, 1
	res := make([][]int, 0)
	res = append(res, intervals[0])//假定第一个是已经合并好的
	//应该拿 res 里len(res)-1 的元素 跟 intervals[i]合并(i>=1 && i<len(intervals))
	for i := 1; i < len(intervals); i++ {
		m := res[len(res)-1]
		n := intervals[i]

		//m[left] 一定是<= n[left]的
		//同一个intervals 里 right 一定 >= left

		//情况一：m[left] <=  m[right] <= n[left]
		if n[left] > m[right] {
			//表示 n 一定跟 m 不相交
			res = append(res, n)
			continue
		}

		//情况二：m[left]  <= n[left]  <=  m[right] <= n[right]
		//if m[right] >= n[left] && m[right] <= n[right] {
		//	m[right] = n[right]  //这里也只需要对 m[right] < n[right] 时有效，m[right] = n[right] 也无需变动
		//}

		//情况三：m[left]  <= n[left]  <=  n[right]  < m[right]
		//if  m[right] > n[right] {
		//	//无需变动
		//}

		if n[right] > m[right] {
			m[right] = n[right]
		}

	}
	return res
}