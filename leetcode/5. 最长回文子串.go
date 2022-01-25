package leetcode

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
	//ansL,ansR记的是回文子串的下标,这两个字符一定是相等的;
	//right，left是子串左右两边的下标，所对应两个字符一定不相等，其实还是比长度
	left, right, ansL, ansR := 0, 0, 0, 0
	// 终止条件为len(s)-1可以应对只有一个字符的情况,直接返回
	// 也减少判断最后一次
	for index := 0; index < len(s)-1; {
		left, right = index, index
		//既可以把left/right置到index两侧，又可以往左右找字符
		for ; left >= 0 && s[left] == s[index]; left-- {
		}
		for ; right < len(s) && s[right] == s[index]; right++ {
		}
		//减少重复判断,直接从右侧开始判断
		index = right
		//左右字符一样持续往两侧移动
		for ; left >= 0 && right < len(s) && s[right] == s[left]; left, right = left-1, right+1 {
		}

		if right-left-2 > ansR-ansL {
			ansL, ansR = left+1, right-1
		}
	}
	return s[ansL : ansR+1]
}

func longestPalindrome2(s string) string {
	begin, end := 0, 0
	for i := range s {
		left, right := expand(s, i, i)
		if right-left > end-begin {
			begin, end = left, right
		}
		left, right = expand(s, i, i+1)
		if right-left > end-begin {
			begin, end = left, right
		}
	}
	return s[begin : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
