package structure

import (
	"errors"
	"fmt"
	"sync"
)

/**
循环链表
*/

type CycleSingleList struct {
	mutex *sync.RWMutex //读写并发控制
	Head  *SingleNode   //头结点，单链表的第一个结点
	Tail  *SingleNode   //尾结点，单链表的最后一个结点
	Size  uint          //单链表长度
}

//初始化循环单链表
func (c *CycleSingleList) Init() {
	c.Size = 0
	c.Head = nil
	c.Tail = nil
	c.mutex = new(sync.RWMutex)
}

//添加结点到尾部的下一个结点或者头部
//中间添加结点(头结点，尾结点]
//index是结点下标，范围是[0,size]
//若index == 0，那么是插入到头结点，此时需要根据size判断要不要修改tail
//若index == size 且size > 0，那么是插入到尾节点的下一个节点
//若 0<index<= size-1 ，那么是插入到（头结点,尾结点]的位置
func (c *CycleSingleList) Insert(node *SingleNode, index uint) (bool, error) {
	if node == nil {
		return false, errors.New("node is nil")
	}
	if index < 0 || index > c.Size {
		return false, errors.New("out of range")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if index == 0 {
		//说明插入的是头结点
		if c.Size == 0 {
			c.Tail = node
			node.NextNode = node
		}
		node.NextNode = c.Head
		c.Tail.NextNode = node
		c.Head = node
		c.Size++
		return true, nil
	}
	if c.Size > 0 && index == c.Size {
		//说明插入到尾元素的下一个
		node.NextNode = c.Head
		c.Tail.NextNode = node
		c.Tail = node
		c.Size++
		return true, nil
	}
	var i uint
	preNode := c.Head
	//获取index的上一个结点
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	node.NextNode = preNode.NextNode
	preNode.NextNode = node
	c.Size++
	return true, nil
}

//根据index删除结点
//index是结点下标，范围是[0,size-1]
func (c *CycleSingleList) Delete(index uint) (bool, error) {
	if index < 0 || index > c.Size-1 || c.Size == 0 {
		return false, errors.New("out of range")
	}
	if index == 0 {
		//说明删除的是头结点
		if c.Size == 1 {
			//说明删除头结点后，没有其他节点了
			c.Size--
			c.Head = nil
			c.Tail = nil
			return true, nil
		}
		//说明删除头结点后，还剩余其他节点
		c.Size--
		c.Head = c.Head.NextNode
		return true, nil
	}
	if index == c.Size-1 && c.Size > 1 {
		//删除的是尾结点
		//找到要删除的index的上一个结点
		preNode := c.Head
		var i uint
		for i = 1; i <= index-1; i++ {
			preNode = preNode.NextNode
		}
		preNode.NextNode = c.Head
		c.Tail = preNode
		c.Size--
		return true, nil
	}
	//删除的是(头结点，尾结点)的范围
	preNode := c.Head
	var i uint
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	preNode.NextNode = preNode.NextNode.NextNode
	c.Size--
	return true, nil
}

//查询
func (c *CycleSingleList) Find(index uint) *SingleNode {
	if index > c.Size-1 || c.Size == 0 || index < 0 {
		return nil
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if index == 0 {
		return c.Head
	}
	//找到index的上一个结点
	preNode := c.Head
	var i uint
	for i = 1; i <= index-1; i++ {
		preNode = preNode.NextNode
	}
	return preNode.NextNode
}

func (c *CycleSingleList) Print() {
	if c == nil || c.Size == 0 {
		fmt.Println("cycle list is nil or empty")
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	fmt.Println("cycle list size is ", c.Size)
	node := c.Head
	for node != nil {
		fmt.Println("data is ", node.Data)
		node = node.NextNode
	}
}
