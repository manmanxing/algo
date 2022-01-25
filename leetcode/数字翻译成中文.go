package leetcode

import (
	"fmt"
	"strings"
)


/**
一个金额大小写转换的，即1024 转成壹千零贰拾肆圆整
解题思路：
逐个读取每个数字，然后对应取中文数字，并加上中文单位，最后把一些需要处理的描述，比如零零，换成零，把零万，换成零等。
*/

func TransferChinese1(num int) string {
	chineseMap:=[]string{"圆整","十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	//对应位上的中文
	chineseNum:=[]string{"零", "壹", "贰","叁","肆","伍","陆","柒","捌","玖"}
	//保存每位数字
	var listNum []int

	for num > 0 {
		listNum = append(listNum, num%10)
		num = num / 10
	}

	n := len(listNum)
	chinese := ""
	//注意这里是倒序的
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	//注意替换顺序
	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}

	return chinese
}

func TransferChinese2()  {

}