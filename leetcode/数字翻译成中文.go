package leetcode

import (
	"fmt"
	"strings"
)

/**
有以下的数字
1987654321
需要翻译成中文：十九亿八千七百六十五万四千三百二十一
 */

func TransferChinese(num int) string{
	//chineseMap:=[]string{"圆整","十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseMap:=[]string{"元整","十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	//chineseNum:=[]string{"零", "壹", "贰","叁","肆","伍","陆","柒","捌","玖"}
	chineseNum:=[]string{"零", "一", "二","三","四","五","六","七","八","九"}
	listNum := []int{}
	for ;num >0;num = num/10{
		listNum= append(listNum, num%10)
	}
	n :=len(listNum)
	chinese :=""
	//注意这里是倒序的
	for i:=n-1; i>=0 ;i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	//注意替换顺序
	for {
		copychinese:=chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		}else {
			chinese= copychinese
		}
	}

	return chinese
}

