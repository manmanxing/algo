package sort

/**
快速排序
平均时间复杂度：O(nlogn)，最坏时间复杂度：O(n^2)，空间复杂度：O(logn)
非稳定排序，原地排序
*/

//下面的算法违背了原地排序原则
func quickSort1(sli []int) []int {
	//先判断是否需要继续进行
	len := len(sli)
	if len <= 1 {
		return sli
	}
	//选择第一个元素作为基准
	base_num := sli[0]
	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	//初始化左右两个切片
	left_sli := []int{}  //小于基准的
	right_sli := []int{} //大于基准的
	for i := 1; i < len; i++ {
		if base_num > sli[i] {
			//放入左边切片
			left_sli = append(left_sli, sli[i])
		} else {
			//放入右边切片
			right_sli = append(right_sli, sli[i])
		}
	}

	//再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	left_sli = quickSort1(left_sli)
	right_sli = quickSort1(right_sli)

	//合并
	left_sli = append(left_sli, base_num)
	return append(left_sli, right_sli...)
}

/**
下面是单向快速排序

原理分析：
j从左向右扫描，a[0,i]表示小于等于pivot的部分，a[i+1,j-1]表示大于pivot的部分，a[j,right]表示未知元素
当a[j]大于pivot时，j继续向前，此时大于pivot的部分就增加一个元素
当a[j]小于等于pivot时，注意i的位置，i的下一个就是大于pivot的元素，将i增加1然后交换a[i]和a[j]
交换后小于等于pivot的部分增加1，j增加1，继续扫描下一个。
而i的下一个元素仍然大于pivot，又回到了先前的状态。
*/
func quickSort2(a []int, left, right int) {
	//如果数据的个小数为1或0则不需要排序
	if left < right {
		//最左边的元素作为中轴元素
		//初始化时
		//小于等于pivot的部分，元素个数为0
		//大于pivot的部分，元素个数也为0
		i, j, pivot := left, left+1, a[left]
		for ; j <= right; {
			if a[j] <= pivot {
				//i增加1后交换
				i++
				a[i], a[j] = a[j], a[i]
				//j继续向前，扫描下一个
				j++
			} else {
				//大于pivot的元素增加一个
				j++
			}
		}
		//a[i]及a[i]以前的都小于等于pivot
		//循环结束后a[i+1]及它以后的都大于pivot
		//所以交换a[left]和a[i],这样我们就将中轴元素放到了适当的位置
		a[left], a[i] = a[i], a[left]
		quickSort2(a, left, i-1)
		quickSort2(a, i+1, right)
	}
}


/**
下面是双向快速排序 方法一

原理分析：
假定最左边的元素为中轴元素
使用两个变量i和j，i指向中轴元素的下一个元素,j指向最后一个元素
我们从前往后找，直到找到一个比中轴元素大的
然后从后往前找，直到找到一个比中轴元素小的，然后交换这两个元素
直到这两个变量交错（i > j）（注意不是相遇 i == j，因为相遇的元素还未和中轴元素比较）。
最后对左半数组和右半数组重复上述操作。
 */
func quickSort3(a []int, left, right int) {
	//如果数据的个小数为1或0则不需要排序
	if left < right {
		//设定中轴元素
		i, j, pivot := left+1, right, a[left]
		//只有在 i<=j 时才会进行循环交换
		//当i == j时，i和j同时指向的元素还没有与中轴元素判断
		for ; i <= j; {
			//从左往右方向
			//如果小于等于中轴元素，i++
			//如果大于中轴元素，i不变，for循环停止
			for ; i <= j && a[i] <= pivot; {
				i++
			}
			//从右往左方向
			//如果大于中轴元素，j--
			//如果小于等于中轴元素，j不变，for循环停止
			for ; i <= j && a[j] > pivot; {
				j--
			}
			//这个if判断条件是为了当上面两个for循环停止时，交换元素，然后改变i，j值使其继续for循环
			//当 i > j 时整个切分过程就应该停止了，不能进行交换操作
			//这里的条件可以改为 i<j,因为这里的 i 永远不会等于 j
			//上面两个for循环就已经将 i=j 的情况给过滤了
			if i <= j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		//当循环结束时，一定有i = j+1, 且i指向的元素大于中轴元素，
		//j指向的元素是最后一个（从左边算起）小于等于中轴的元素
		//因此只要将j指向的元素与中轴元素交换
		a[left], a[j] = a[j], a[left] //将中轴元素和j所指的元素互换
		quickSort3(a, left, j-1)      //递归左半部分
		quickSort3(a, j+1, right)     //递归右半部分
	}
}

/**
下面是双向快速排序 方法二

原理分析：
使用两个变量i和j，i指向最左边的元素，j指向最右边的元素
将首元素作为中轴，将首元素复制到变量pivot中
这时可以将首元素i所在的位置看成一个坑
从j的位置从右向左扫描，找一个小于等于中轴的元素a[j]，来填补a[i]这个坑
填补完成后，拿去填坑的元素所在的位置j又可以看做一个坑
这时以i的位置从前往后找一个大于中轴的元素来填补a[j]这个新的坑
如此往复，直到i和j相遇（i == j，此时i和j指向同一个坑）。
最后将中轴元素放到这个坑中。最后对左半数组和右半数组重复上述操作。
 */
func quickSort4(a []int, left, right int) {
	//如果数据的个小数为1或0则不需要排序
	if left < right {
		//最左边的元素作为中轴复制到 pivot，这时最左边的元素可以看做一个坑
		//注意这里 i = left,而不是 i = left+1, 因为i代表坑的位置,当前坑的位置位于最左边
		i, j, pivot := left, right, a[left]
		for ; i < j; {
			//下面面两个循环的位置不能颠倒，因为第一次坑的位置在最左边
			for ; i < j && a[j] > pivot; {
				j--
			}
			//填a[i]这个坑,填完后a[j]是个坑
			//注意不能是a[i++] = a[j]
			//当因i==j时跳出上面的循环时,坑为i和j共同指向的位置
			//执行a[i++] = a[j],会导致i比j大1,但此时i并不能表示坑的位置
			a[i] = a[j]

			for ; i < j && a[i] <= pivot; {
				i++
			}
			//填a[j]这个坑，填完后a[i]是个坑
			//同理不能是a[j--] = a[i]
			a[j] = a[i]
		}
		//循环结束后i和j相等，都指向坑的位置，将中轴填入到这个位置
		a[i] = pivot
		quickSort4(a, left, i-1)//递归左边的数组
		quickSort4(a, i+1, right)//递归右边的数组
	}
}