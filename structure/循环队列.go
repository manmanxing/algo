package structure

import (
	"errors"
	"fmt"
)

/**
循环队列
使用线性表的顺序存储结构
判断队列是空还是满
1.设置一个标志变量flag
2.使用一个变量记录入栈时就加1, 出栈时就减1
3.保留一个元素空间，不作存储使用
本示例将使用第三种方法
*/

//循环队列存储的数据类型
type CycleQueueData interface{}

//循环队列
type CycleQueue struct {
	data  []CycleQueueData //循环队列存储的数据
	Front int              //头结点的下标，负责弹出数据
	Tail  int              //负责压入数据，若队列不空，则指向队列尾结点的下一个位置
	Size  int              //存储的数据长度
}

//初始化循环队列
func (c *CycleQueue) Init(cap int) {
	c.Size = cap
	c.data = make([]CycleQueueData, c.Size)
	c.Front = 0
	c.Tail = 0
}

//计算队列长度
func (c *CycleQueue) QueueLength() int {
	return (c.Tail - c.Front + c.Size) % c.Size
}

//入队操作
func (c *CycleQueue) Push(e CycleQueueData) (bool, error) {
	//如果队列已经满了
	if c.IsFull() {
		return false, errors.New("queue is full")
	}
	//尾部添加数据
	c.data[c.Tail] = e
	//尾部元素指向下一个空间位置,取模运算保证了索引不越界（余数一定小于除数）
	//若到最后则转到数组头部
	c.Tail = (c.Tail + 1) % c.Size
	return true, nil
}

//出队操作
func (c *CycleQueue) Pop() CycleQueueData {
	//如果队为空
	if c.IsEmpty() {
		return nil
	}
	head := c.data[c.Front]
	//清空下标为c.Front的数据
	c.data[c.Front] = nil
	//Front 向后移动一位，若已是最后的位置则转移到数组头部
	c.Front = (c.Front + 1) % c.Size
	return head
}

//判断队列是否已满
func (c *CycleQueue) IsFull() bool {
	if (c.Tail+1)%c.Size == c.Front {
		return true
	}
	return false
}

//判断队列是否空
func (c *CycleQueue) IsEmpty() bool {
	if c.Front == c.Tail {
		return true
	}
	return false
}

//打印循环列表
func (c *CycleQueue) Print() {
	fmt.Println("queue:", c.data)
}
