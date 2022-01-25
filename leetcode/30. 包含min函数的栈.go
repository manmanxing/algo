package leetcode

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
*/

type MinStack struct {
	StackA []int // 原stack
	StackB []int // 记录StackA每个时期的最小值
}

func ConstructorMinStack() MinStack {
	return MinStack{
		StackA: nil,
		StackB: nil,
	}
}

func (this *MinStack) pushB(min int) {
	this.StackB = append(this.StackB, min)
}

func (this *MinStack) Push(x int) {
	if len(this.StackA) <= 0 {
		// 第一次
		this.StackA = append(this.StackA, x)
		this.pushB(x)
		return
	}
	// 不是第一次，拿 x 与 stackB的最后一个元素比较大小
	this.StackA = append(this.StackA, x)
	lastB := this.StackB[len(this.StackB)-1]
	if x >= lastB {
		this.pushB(lastB)
	} else {
		this.pushB(x)
	}
}

// 输出最后一个原元素
func (this *MinStack) Pop() {
	lastA := len(this.StackA) - 1
	lastB := len(this.StackB) - 1
	this.StackA = this.StackA[:lastA]
	// 同时清除stackB最后一个元素
	this.StackB = this.StackB[:lastB]
}

// 打印最后一个元素
func (this *MinStack) Top() int {
	lastA := len(this.StackA) - 1
	return this.StackA[lastA]
}

func (this *MinStack) Min() int {
	return this.StackB[len(this.StackB)-1]
}
