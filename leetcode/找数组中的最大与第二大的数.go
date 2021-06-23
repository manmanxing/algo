package leetcode


/**
示例
[1,2,3,4,5]，结果为4，5
[1,1,1],结果为1，1
[0],结果为0，0
 */

func getMaxAndSecondNum(a []int)(max,sMax int)  {
	if len(a) <= 0 {
		return 0,0
	}

	for i := range a {
		if max < a[i] {
			sMax = max
			max = a[i]
		}
	}

	return
}