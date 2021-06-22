package sort

//归并排序
//时间复杂度：O(nlogn) 空间复杂度：O(n)
//稳定排序，非原地排序
//b:切片长度，a：等于0，起始下标
func mergerSort(arr []int, a, b int) {
	//如果切片只有一个元素，则不需要递归排序
	if b-a <= 1 {
		return
	}
	//将切片分为两个切片
	c := (a + b) / 2
	//对左半部分进行归并排序
	mergerSort(arr, a, c)
	//对右半部分进行归并排序
	mergerSort(arr, c, b)

	//创建左右部分的切片，并copy元素
	arrLeft := make([]int, c-a)
	arrRight := make([]int, b-c)
	copy(arrLeft, arr[a:c])
	copy(arrRight, arr[c:b])
	i := 0//用于左边切片
	j := 0//用于右边切片
	//遍历切片所有元素，将左右数组合并
	for k := a; k < b; k++ {
		if i >= c-a {
			arr[k] = arrRight[j]
			j++
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] < arrRight[j] { //决定是递增还是递减，小切片时调用
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}
