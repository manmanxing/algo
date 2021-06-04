package structure

import (
	"errors"
	"fmt"
	"sync"
)

/**
线性表的链式存储结构：单链表
*/

//结点数据
//SingleItem 可以理解为范型，也就是任意的数据类型
type SingleItem interface{}

//单链表结点
//结点除了自身的数据外，还必须有指向下一个结点的地址域
type SingleNode struct {
	//数据域
	Data SingleItem
	//地址域
	NextNode *SingleNode
}

//单链表
type SingleList struct {
	mutex *sync.RWMutex //读写并发控制
	Head  *SingleNode   //头结点，单链表的第一个结点
	Tail  *SingleNode   //尾结点，单链表的最后一个结点
	Size  uint          //单链表长度
}

//初始化单链表
func (list *SingleList) Init() {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	list.mutex = new(sync.RWMutex)
}

//添加结点到尾部的下一个结点或者头部
//中间添加结点(头结点，尾结点]
//index是结点下标，范围是[0,size]
//若index == 0，那么是插入到头结点，此时需要根据size判断要不要修改tail
//若index == size 且size > 0，那么是插入到尾节点的下一个节点
//若 0<index<= size-1 ，那么是插入到（头结点,尾结点]的位置
func (list *SingleList) Insert(node *SingleNode, index uint) (bool, error) {
	if node == nil {
		return false, errors.New("node is nil")
	}
	if index > list.Size || index < 0 {
		//是否越界
		//index == list.size 时表示插入到尾结点的下一个结点
		return false, errors.New("out of range")
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		//说明插入的是头结点
		if list.Size == 0 {
			list.Tail = node
		}
		node.NextNode = list.Head
		list.Head = node
		list.Size++
		return true, nil
	}
	if index == list.Size && list.Size > 0 {
		//说明插入的是尾结点的下一个结点
		list.Tail.NextNode = node
		list.Tail = node
		list.Size++
		return true, nil
	}
	//说明插入的是（头结点,尾结点]的位置
	var i uint
	preNode := list.Head
	//获取index的上一个结点
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	next := preNode.NextNode
	node.NextNode = next
	preNode.NextNode = node
	list.Size++
	return true, nil
}

//根据index删除结点
//index是结点下标，范围是[0,size-1]
func (list *SingleList) Delete(index uint) (bool, error) {
	if list == nil || list.Size == 0 || index > list.Size-1 || index < 0 {
		return false, nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		//删除的是头结点
		//如果list的长度本来为1，那么表示list被删除
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
			list.Size--
			return true, nil
		}
		//说明删除头结点后，list还剩余结点
		list.Head = list.Head.NextNode
		list.Size--
		return true, nil
	}
	if index == list.Size-1 && list.Size > 1 {
		//删除的是尾结点
		//找到要删除的index的上一个结点
		preNode := list.Head
		var i uint
		for i = 1; i <= index-1; i++ {
			preNode = preNode.NextNode
		}
		preNode.NextNode = nil
		list.Tail = preNode
		list.Size--
		return true, nil
	}
	//删除的是(头结点，尾结点)的范围
	//找到要删除的index的上一个结点
	preNode := list.Head
	var i uint
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	preNode.NextNode = preNode.NextNode.NextNode
	list.Size--
	return true, nil
}

//根据index查询结点
//index是结点下标，范围是[0,size-1]
func (list *SingleList) Find(index uint) *SingleNode {
	if list == nil || index > list.Size-1 || list.Size == 0 || index < 0 {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		return list.Head
	}
	//找到index的上一个结点
	preNode := list.Head
	var i uint
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	return preNode.NextNode
}

func (list *SingleList) Print() {
	if list == nil || list.Size == 0 {
		fmt.Println("singleList is nil or empty")
		return
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	fmt.Println("singleList size is ", list.Size)
	node := list.Head
	for node != nil {
		fmt.Println("data is ", node.Data)
		node = node.NextNode
	}
}

//单链表反转：示例一
//空间复杂度是 O(1),时间复杂度为O(n)
func Reverse(head *SingleNode) *SingleNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *SingleNode = nil
	for cur != nil {
		pre, cur, cur.NextNode = cur, cur.NextNode, pre
	}
	return pre
}

//单链表反转：示例二
//借助函数调用栈的思想，其实本质上也是一个栈。
//参数是需要反转的单链表的头节点
//返回的是单链表反转后的头节点
func reverseByRecursion(head *SingleNode) *SingleNode {
	if head == nil || head.NextNode == nil {
		return head
	}
	newHead := reverseByRecursion(head.NextNode)
	head.NextNode.NextNode = head
	head.NextNode = nil
	return newHead
}

//检测单链表是否有环
func HasCycle(head *SingleNode) bool {
	if nil != head {
		slow := head
		fast := head
		for nil != fast && nil != fast.NextNode {
			slow = slow.NextNode
			fast = fast.NextNode.NextNode
			if slow == fast {
				return true
			}
		}
	}
	return false
}

//求单链表环的长度
func CycleLen(head *SingleNode) int {
	len := 0
	if nil != head {
		slow := head
		fast := head
		for nil != fast && nil != fast.NextNode {
			slow = slow.NextNode
			fast = fast.NextNode.NextNode
			len++
			if slow == fast {
				break
			}
		}
	}
	return len
}

//两个有序单链表的合并
//将两个有序链表合并为一个新的有序链表并返回
func mergeSortedList(l1, l2 *SingleNode) *SingleNode {
	var res *SingleNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	//使用递归，不断去找两个链表中比较小的元素，然后result接上那个元素
	if l1.Data.(int) >= l2.Data.(int) {
		res = l2
		res.NextNode = mergeSortedList(l1, l2.NextNode)
	} else {
		res = l1
		res.NextNode = mergeSortedList(l1.NextNode, l2)
	}
	return res
}

//获取单链表中间结点
func FindMiddleNode(head *SingleNode) *SingleNode {
	if nil == head || nil == head.NextNode {
		return nil
	}
	if nil == head.NextNode.NextNode {
		return head.NextNode
	}

	slow, fast := head, head
	for nil != fast && nil != fast.NextNode {
		//fast是等比数列，比值为2
		//等fast移动到最后一个值或者超过最后一个值，也就是for循环条件不成立
		//此时slow刚刚好移动到中间
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
	}
	return slow
}
