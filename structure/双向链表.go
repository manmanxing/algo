package structure

import (
	"errors"
	"fmt"
	"sync"
)

/**
双向链表
*/

//双向链表结点
//结点除了自身的数据外，还必须有指向下一个结点的地址域和上一个结点的地址域
type DoubleNode struct {
	//数据域
	Key, Value interface{}
	//上一个地址域
	PrevNode *DoubleNode
	//下一个地址域
	NextNode *DoubleNode
}

//双链表
type DoubleList struct {
	mutex *sync.Mutex //并发控制
	Head  *DoubleNode //头结点，双链表的第一个结点
	Tail  *DoubleNode //尾结点，双链表的最后一个结点
	Size  int         //双链表长度
}

func InitDoubleNode(k, v interface{}) *DoubleNode {
	return &DoubleNode{
		Key:   k,
		Value: v,
	}
}

func (node *DoubleNode) String() string {
	return fmt.Sprintf("{%v:%v}", node.Key, node.Value)
}

//初始化双链表
func InitDoubleList(size int) (list *DoubleList, err error) {
	if size < 0 {
		return nil, errors.New("size < 0")
	}
	list = &DoubleList{
		mutex: new(sync.Mutex),
		Head:  nil,
		Tail:  nil,
		Size:  size,
	}
	return
}

//打印双向链表
func (list *DoubleList) String() string {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	result := ""
	if list == nil || list.Size == 0 {
		result = "doubleList is nil or empty"
		return result
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	node := list.Head
	for node != nil {
		result += node.String()
		if node.NextNode != nil {
			result += "=>"
		}
		node = node.NextNode
	}
	return result
}

//在双向链表中的头结点后添加节点
func (list *DoubleList) AddHead(node *DoubleNode) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if node == nil {
		return
	}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		list.Head.PrevNode = nil
		list.Tail.NextNode = nil
		list.Size++
		return
	}

	list.Head.PrevNode = node
	node.NextNode = list.Head
	list.Head = node
	list.Head.PrevNode = nil
	list.Size++
	return
}

//在双向链表中的尾节点后添加节点
func (list *DoubleList) AddTail(node *DoubleNode) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if node == nil {
		return
	}
	if list.Tail == nil {
		list.Head = node
		list.Tail = node
		list.Head.PrevNode = nil
		list.Tail.NextNode = nil
		list.Size++
		return
	}

	list.Tail.NextNode = node
	node.PrevNode = list.Tail
	list.Tail = node
	list.Tail.NextNode = nil
	list.Size++
	return
}

//删除头节点
func (list *DoubleList) RemoveHead() {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Head == nil {
		return
	}

	if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
		list.Size--
		return
	}
	next := list.Head.NextNode
	next.PrevNode = nil
	list.Head = next
	list.Size--
	return
}

//删除尾节点
func (list *DoubleList) RemoveTail() {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Tail == nil {
		return
	}

	if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
		list.Size--
		return
	}

	pre := list.Tail.PrevNode
	pre.NextNode = nil
	list.Tail = pre
	list.Size--
}

//删除某一个双向链表中的节点
func (list *DoubleList) Remove(node *DoubleNode) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Size <= 0 || node == nil {
		return
	}

	if list.Head == node {
		list.RemoveHead()
		return
	}

	if list.Tail == node {
		list.RemoveTail()
		return
	}

	pre := node.PrevNode
	next := node.NextNode
	pre.NextNode = next
	next.PrevNode = pre
	list.Size--
	return
}

//弹出头节点
func (list *DoubleList) Pop() *DoubleNode {
	head := list.Head
	list.RemoveHead()
	return head
}
