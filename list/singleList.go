package list

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
//默认插入时，index后面的元素全部往后移
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
	if list.Size == 0 {
		//说明插入的是头结点
		node.NextNode = list.Head
		list.Head = node
		list.Tail = node
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
//默认删除时，所有index后面的元素都往前移
func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 || index < 0 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		//删除的是头结点
		head := list.Head.NextNode
		list.Head = head
		//如果list的长度本来为1，那么表示list被删除
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size--
		return true
	}
	if index == list.Size-1 {
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
		return true
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
	return true
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
