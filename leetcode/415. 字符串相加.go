package leetcode

/**
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：
num1 和num2的长度都小于 5100
num1 和num2 都只包含数字0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式

解题思路：
大数相加法
*/

func addStrings(num1 string, num2 string) string {
	//将字符串转换为byte数组，并且保证 nb1 长度 >= nb2 长度
	nb1, nb2 := []byte(num1), []byte(num2)
	if len(nb1) < len(nb2) {
		nb1, nb2 = nb2, nb1
	}

	//sum 表示该位上的和，比如 1+1 = 2，8+9 = 17
	//"0" 的byte 是 48
	sum := byte(0)
	for i, j := len(nb1)-1, len(nb2)-1; i >= 0; {
		if j >= 0 {
			sum += nb2[j] - '0'
			j--
		}
		sum += nb1[i] - '0'
		//让sum 对 10 取余，然后 + "0"，表示byte位
		nb1[i] = (sum % 10) + '0'
		//会每次让 sum /10，表示是否需要进位，1表示需要，0表示不需要
		i, sum = i-1, sum/10
	}
	//判断最终是否还需要进位
	if sum != 0 {
		nb1 = append([]byte{'1'}, nb1...)
	}
	return string(nb1)
}
