package sort

import "fmt"

//先构建小顶堆
//再进行排序：首先将堆顶的最小值与最后一个值交换（数组中的0和11位置交换）
//交换后重新构建堆（11位置已经是最小的值，不参与堆的构建，所以重新构建堆的数组要减去最后的值）
//如此往返，排序就完成
func minHeap(root int, end int, c []int)  {
	for {
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}
		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

// HeapSort 降序排序
func HeapSort(c []int)  {
	//第一个需要比较的root节点为(len(arr)-1)/2
	//然后一直递减到根结点
	var n = len(c)-1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	fmt.Println("堆构建完成")
	for end := n; end >=0; end-- {
		if c[0]<c[end]{
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
}