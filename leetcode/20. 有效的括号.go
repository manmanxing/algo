package leetcode

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例一
输入：s = "()"
输出：true

示例二
输入：s = "()[]{}"
输出：true

示例三
输入：s = "(]"
输出：false

示例四
输入：s = "([)]"
输出：false

示例五
输入：s = "{[]}"
输出：true

解题思路：
比如 "()[]{}", 有 "),},]" 表示一定会有 "(,{,[" 与之对应，否则就不符合题意，而且在遇到 "),},]" 中的任意一个时，与之相邻的左边的那个也一定是 "(,{,[" 与之对应。
*/


func IsValid1(s string) bool {
	var stack []string

	for _, ch := range s {
		c := string(ch)
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == "(" && c == ")" || top == "[" && c == "]" || top == "{" && c == "}" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

