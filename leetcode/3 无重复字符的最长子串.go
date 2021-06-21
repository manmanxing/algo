package leetcode

import "strings"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例1
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例2
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例3
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

示例4
输入: s = ""
输出: 0

解题思路1：
窗口可以在两个边界 start,end 移动，一开始窗口大小为0
1.窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
2.窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符的后一位
3.遍历结束，返回end-start，窗口大小

解题思路2：
窗口可以在两个边界 start,end 移动，一开始窗口大小为0
随着窗口的前进，窗口右侧 end 值会依次增大
每次查询窗口里的字符，若窗口中有查询的字符，窗口的左侧移动到该字符加一的位置。
只有在每次查询窗口中没有重复的字符时，每次记录窗口的最大程度，重复操作直到数组遍历完成
*/

func lengthOfLongestSubstring1(s string) int {
	if len(strings.TrimSpace(s)) <= 0 {
		return 0
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}

	return end - start
}

func lengthOfLongestSubstring2(s string) int {
	if len(strings.TrimSpace(s)) <= 0 {
		return 0
	}

	start, end, length := 0, 0, 0
	lenStr := len(s)
	for i := 0; i < lenStr; i++ {
		demo1 := s[start:end]
		demo2 := string(s[i])
		index := strings.Index(demo1, demo2)
		if index == -1 {
			end++
			if end-start > length {
				length = end - start
			}
		} else {
			end++
			start += index + 1
		}
	}

	return length
}
