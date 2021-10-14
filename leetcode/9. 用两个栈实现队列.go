package leetcode

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

示例一
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

示例二
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]

解题思路：
题目的意思是：栈的特性是先进后出，而队列的特性是先进先出，利用栈来实现队列的这个特性。
一个栈存入，另一个栈取出
*/

type CQueue struct {
	StackA []int //负责插入
	StackB []int //负责取出
}

func ConstructorCQueue() CQueue {
	return CQueue{
		StackA: nil,
		StackB: nil,
	}
}

//负责插入
func (this *CQueue) AppendTail(value int) {
	if value > 10000 || value < 1 {
		return
	}

	this.StackA = append(this.StackA, value)
}

//负责取出
func (this *CQueue) DeleteHead() int {
	result := -1
	//如果 stackB 没有元素则从 stackA 中取出所有
	if len(this.StackB) <= 0 {
		//如果A没有元素
		if len(this.StackA) <= 0 {
			return result
		}

		//将A中所有的元素都放入到B中
		for len(this.StackA) > 0 {
			last := len(this.StackA)-1
			this.StackB = append(this.StackB, this.StackA[last])
			this.StackA = this.StackA[:last]
		}
	}

	last := len(this.StackB) -1
	result = this.StackB[last]
	this.StackB = this.StackB[:last]
	return result
}
