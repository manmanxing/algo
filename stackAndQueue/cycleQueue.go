package stackAndQueue

/**
循环队列
使用单链表结构
判断队列是空还是满
1.设置一个标志变量flag
2.使用一个变量记录入栈时就加1, 出栈时就减1
3.保留一个元素空间，不作存储使用
本示例将使用第三种方法
*/

type CycleQueueData interface{}

type CycleQueue struct {
	data  []CycleQueueData //循环队列存储的数据
	Front int              //头结点，负责弹出数据
	Tail  int              //尾结点，负责压入数据
	Size  int              //存储的数据长度
}
