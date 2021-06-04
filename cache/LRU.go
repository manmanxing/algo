package cache

import (
	"errors"
	"sync"

	"github.com/manmanxing/algo/structure"
)

type LRUCache struct {
	Size     int //lru的长度 , Size <= Capacity
	Capacity int //lru的容量 ,初始化时，此值必须 >0
	Head     *structure.DoubleNode
	Tail     *structure.DoubleNode
	mutex    *sync.Mutex                           //读写并发控制
	Nodes    map[interface{}]*structure.DoubleNode //map的key就是DoubleNode里的key，实现查找结点时间复杂度为 O(1)
}

//初始化
func InitLRU(capacity int) (lru *LRUCache, err error) {
	if capacity <= 0 {
		return nil, errors.New("capacity <= 0")
	}
	lru = new(LRUCache)
	lru.Nodes = make(map[interface{}]*structure.DoubleNode, capacity)
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

	if key == nil {
		return -1
	}

	if v, ok := lru.Nodes[key]; ok {
		lru.removeNodeToHead(v)
		return v
	}
	return -1
}

//新增
//会直接插入到链表头部
//需要判断容量是否够用，不够就移除掉尾结点再插入新结点
func (lru *LRUCache) Add(key, value interface{}) {
	lru.mutex.Lock()
	defer func() {
		lru.mutex.Unlock()
	}()

	if key == nil || value == nil {
		return
	}

	newNode := structure.InitDoubleNode(key, value)

	//如果是空的双向链表
	if lru.Tail == nil && lru.Head == nil {
		lru.Tail = newNode
		lru.Head = newNode
		lru.Nodes[key] = newNode
		lru.Size++
		return
	}

	//如果本身就存在，直接移到head
	if v, ok := lru.Nodes[key]; ok {
		lru.removeNodeToHead(v)
		return
	} else {
		//不存在
		//首先判断是否超容
		if lru.Size >= lru.Capacity {
			//说明map存满了
			lru.removeLastNode()
		}
		//插入到头部
		lru.Head.PrevNode = newNode
		newNode.NextNode = lru.Head
		newNode.PrevNode = nil
		lru.Head = newNode
		lru.Nodes[key] = newNode
		lru.Size++
		return
	}
}

//将指定的结点移向表头
func (lru *LRUCache) removeNodeToHead(node *structure.DoubleNode) {
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
		delete(lru.Nodes, lru.Tail.Key)
		lru.Head = nil
		lru.Tail = nil
		lru.Size--
		return
	}

	delete(lru.Nodes, lru.Tail.Key)
	lru.Tail = lru.Tail.PrevNode
	lru.Tail.NextNode = nil
	lru.Size--
}
