package leetcode

import "strings"

/**
给你一个字符串 s，找到 s 中最长的回文子串。

示例1
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例2
输入：s = "cbbd"
输出："bb"

示例3
输入：s = "a"
输出："a"

示例4
输入：s = "ac"
输出："a"

解题思路
中心扩散法：
例如： 字符串abcba 共有5（字母） + 4（两字母间） = 9个中心点；
因此，长度为N的string共有2N-1个中心。我们的目标就是统计以这2N-1个点为中心的最长回文串s1,s2,..,s2N-1，并从中挑出全局最长回文串。
保留最大长度回文串index，记为left和right；完成遍历后返回以left和right为边界的substring



*/

func longestPalindrome1(s string) string {
	if len(strings.TrimSpace(s)) <= 1 {
		return s
	}

	return ""
}
