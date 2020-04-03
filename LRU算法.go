package main

import "errors"

//使用双向链表实现
//使用 map，实现查找结点时间复杂度为 O(1)
//假设双向结点中保存的值为 整型

type LRUCache struct {
	Size  int //最大长度,初始化时，此值必须 >0
	Nodes map[int]*DoubleNode
	Head  *DoubleNode
	Tail  *DoubleNode
}

//初始化
func (lru *LRUCache) Init(capacity int) {
	lru.Nodes = make(map[int]*DoubleNode, capacity)
	lru.Size = capacity
	lru.Head = nil
	lru.Tail = nil
}

//查询
//如果密钥 (key) 存在于缓存中，则获取密钥的值,然后将查询的结点放置表头，否则返回 -1
func (lru *LRUCache) Find(key int) int {
	if v, ok := lru.Nodes[key]; ok {
		lru.removeNodeToFirst(v)
		return v.Data.(int)
	}
	return -1
}

//新增
//会直接插入到链表头部
//需要判断容量是否够用，不够就移除掉尾结点再插入新结点
func (lru *LRUCache) Insert(k int, value int) error {
	if lru.Size <= 0 {
		return errors.New("lru size <= 0")
	}
	node := &DoubleNode{
		Data:     value,
		PrevNode: nil,
		NextNode: nil,
	}
	if lru.Tail == nil && lru.Head == nil {
		//如果是空的双向链表
		lru.Tail = node
		lru.Head = node
		lru.Nodes[k] = node
		return nil
	}
	//如果不是空的双向链表
	if len(lru.Nodes) == lru.Size {
		//说明map存满了
		lru.removeLastNode()
	}
	lru.Head.PrevNode = node
	node.NextNode = lru.Head
	lru.Head = node
	lru.Nodes[k] = node
	return nil
}

//将指定的结点移向表头
func (lru *LRUCache) removeNodeToFirst(node *DoubleNode) {
	node.PrevNode = nil
	node.NextNode = lru.Head
	lru.Head.PrevNode = node
}

//移除最后有一个结点
//需要判断链表的长度
func (lru *LRUCache) removeLastNode() {
	if len(lru.Nodes) == 1 {
		delete(lru.Nodes, lru.Tail.Data.(int))
		lru.Head = nil
		lru.Tail = nil
		return
	}
	delete(lru.Nodes, lru.Tail.Data.(int))
	lru.Tail = lru.Tail.PrevNode
	lru.Tail.NextNode = nil
}
