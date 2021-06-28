package leetcode

/**
定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例一
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]

示例二
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12],
  [13,14,15,16]
]
输出: [1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10]

解题思路：
运动方向会从右到下，再到左，再到上，这样的规律进行。
每次运动完后，边界都会发生变化
1、向右运动结束后，上边界下移，即top++
2、向下运动结束后，右边界左移，即right--
3、向左运动结束后，下边界上移，即bottom--
4、向上运动结束后，左边界右移，即left++

每次运动结束后，开始新方向运动时，对应的开始行列也会发生变化
1、向右运动结束后，开始行下移，即row++
2、向下运动结束后，开始列左移，即col--
3、向左运动结束后，开始行上移，即row--
4、向上运动结束后，开始列右移，即col++

运动结束的时候，也就是 上下边界 和 左右边界 重叠的时候。
*/

func spiralOrder(matrix [][]int) []int {
	var result []int

	if len(matrix) <= 0 {
		return result
	}

	//dir 表示 1、向右 2、向下 3、向左 4、向上 的方向
	//默认是向右开始
	dir := 1
	//行与列都是从0开始
	row, col := 0, 0

	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0

	for top <= bottom && left <= right {
		//将对应行对应列的数据装进来
		result = append(result, matrix[row][col])

		//选择走向
		switch dir {
		case 1:
			//向右
			if col == right {
				//已经到了最右边，开始向下
				dir = 2
				top++
				row++
				continue
			}
			//还没有到最右边
			col++
		case 2:
			//向下
			if row == bottom {
				//已经到了最下边，开始向左
				dir = 3
				right--
				col--
				continue
			}
			row++
		case 3:
			//向左
			if col == left {
				//已经到了最左边，开始向上
				dir = 4
				bottom--
				row--
				continue
			}
			col--
		case 4:
			//向上
			if row == top {
				//已经到了最上边，开始向右
				dir = 1
				left++
				col++
				continue
			}
			row--
		}
	}
	return result
}