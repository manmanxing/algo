package main

import (
	"errors"
	"sync"
)

type LRUCache struct {
	Size     int //lru的长度 , Size <= Capacity
	Capacity int //lru的容量 ,初始化时，此值必须 >0
	Head     *DoubleNode
	Tail     *DoubleNode
	mutex    *sync.Mutex                 //读写并发控制
	Nodes    map[interface{}]*DoubleNode //map的key就是DoubleNode里的key，实现查找结点时间复杂度为 O(1)
}

//初始化
func InitLRU(capacity int) (lru *LRUCache, err error) {
	if capacity <= 0 {
		return nil, errors.New("capacity <= 0")
	}
	lru = new(LRUCache)
	lru.Nodes = make(map[interface{}]*DoubleNode, capacity)
	lru.Size = 0
	lru.Capacity = capacity
	//懒加载
	lru.Head = nil
	lru.Tail = nil
	return
}

//查询
//如果密钥 (key) 存在于缓存中，则获取密钥的值,然后将查询的结点放置表头，否则返回 -1
func (lru *LRUCache) Find(key interface{}) interface{} {
	lru.mutex.Lock()
	defer func() {
		lru.mutex.Unlock()
	}()

	if v, ok := lru.Nodes[key]; ok {
		lru.removeNodeToFirst(v)
		return v
	}
	return -1
}

//新增
//会直接插入到链表头部
//需要判断容量是否够用，不够就移除掉尾结点再插入新结点
func (lru *LRUCache) Insert(k interface{}, value *DoubleNode) {
	lru.mutex.Lock()
	defer func() {
		lru.mutex.Unlock()
	}()

	//首先判断是否超容
	if lru.Size >= lru.Capacity {
		//说明map存满了
		lru.removeLastNode()
	}
	//如果是空的双向链表
	if lru.Tail == nil && lru.Head == nil {
		lru.Tail = value
		lru.Head = value
		lru.Nodes[k] = value
		lru.Size++
		return
	}

	//插入到头部
	lru.Head.PrevNode = value
	value.NextNode = lru.Head
	value.PrevNode = nil
	lru.Head = value
	lru.Nodes[k] = value
	lru.Size ++
	return
}

//将指定的结点移向表头
func (lru *LRUCache) removeNodeToFirst(node *DoubleNode) {
	preNode := node.PrevNode
	preNode.NextNode = node.NextNode
	node.NextNode.PrevNode = preNode
	node.PrevNode = nil
	node.NextNode = lru.Head
	lru.Head.PrevNode = node
	lru.Head = node
}

//移除最后一个node
func (lru *LRUCache) removeLastNode() {

	if lru.Size <= 0 {
		return
	}
	if lru.Size == 1 {
		delete(lru.Nodes,lru.Tail.Key)
		lru.Head = nil
		lru.Tail = nil
		return
	}

	delete(lru.Nodes, lru.Tail.Key)
	lru.Tail = lru.Tail.PrevNode
	lru.Tail.NextNode = nil
}