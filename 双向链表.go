package main

import (
	"errors"
	"fmt"
	"sync"
)

/**
双向链表
*/

//结点数据
//DoubleItem 可以理解为范型，也就是任意的数据类型
type DoubleItem interface{}

//双向链表结点
//结点除了自身的数据外，还必须有指向下一个结点的地址域和上一个结点的地址域
type DoubleNode struct {
	//数据域
	Data DoubleItem
	//上一个地址域
	PrevNode *DoubleNode
	//下一个地址域
	NextNode *DoubleNode
}

//双链表
type DoubleList struct {
	mutex *sync.RWMutex //读写并发控制
	Head  *DoubleNode   //头结点，双链表的第一个结点
	Tail  *DoubleNode   //尾结点，双链表的最后一个结点
	Size  uint          //双链表长度
}

//初始化双链表
func (list *DoubleList) Init() {
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
func (list *DoubleList) Insert(index uint, node *DoubleNode) (bool, error) {
	if node == nil {
		return false, errors.New("node is nil")
	}
	if index > list.Size || index < 0 {
		//说明越界
		return false, errors.New("out of range")
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		//说明插入的是头结点
		if list.Size == 0 {
			//此时list没有结点
			list.Tail = node
		}
		node.NextNode = list.Head
		list.Head = node
		list.Head.PrevNode = nil
		list.Size++
		return true, nil
	}
	if index == list.Size && list.Size > 0 {
		//说明是插入到尾结点的下一个结点
		node.PrevNode = list.Tail
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
	node.PrevNode = preNode
	node.NextNode = next
	preNode.NextNode = node
	next.PrevNode = node
	list.Size++
	return true, nil
}

//根据index删除结点
//index是结点下标，范围是[0,size-1]
func (list *DoubleList) Delete(index uint) (bool, error) {
	if index > list.Size-1 || index < 0 {
		return false, errors.New("out of range")
	}
	if list == nil {
		return false, errors.New("DoubleList is nil")
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		//说明删除的是头结点
		if list.Size == 1 {
			//此时删除后，list 没有结点
			list.Head = nil
			list.Tail = nil
			list.Size--
			return true, nil
		}
		//说明删除头结点后，list 还剩余结点
		list.Head = list.Head.NextNode
		list.Head.PrevNode = nil
		list.Size--
		return true, nil
	}
	if index == list.Size-1 && list.Size > 1 {
		//说明删除的是尾节点
		list.Tail = list.Tail.PrevNode
		list.Tail.NextNode = nil
		list.Size--
		return true, nil
	}
	//说明删除的是中间结点(头结点，尾结点)
	var i uint
	node := list.Head
	//找到要删除的结点的上一个结点
	for i = 1; i <= index-1; i++ {
		node = node.NextNode
	}
	node.NextNode = node.NextNode.NextNode
	node.NextNode.PrevNode = node
	return true, nil
}

//查询结点
//index是结点下标，范围是[0,size-1]
func (list *DoubleList) Find(index uint) *DoubleNode {
	if list == nil || index > list.Size-1 || index < 0 {
		return nil
	}
	if index == 0 {
		//查询的是头结点
		return list.Head
	}
	var i uint
	node := list.Head
	//查询结点的上一个结点
	for i = 1; i <= index-1; i++ {
		node = node.NextNode
	}
	return node.NextNode
}

//打印双向链表
func (list *DoubleList) Print() {
	if list == nil || list.Size == 0 {
		fmt.Println("doubleList is nil or empty")
		return
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	fmt.Println("doubleList size is ", list.Size)
	node := list.Head
	for node != nil {
		fmt.Println("data is ", node.Data)
		node = node.NextNode
	}
}
