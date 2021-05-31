package main

import (
	"errors"
	"sync"
)

/**
在FIFO Cache设计中，核心原则就是：如果一个数据最先进入缓存中，则应该最早淘汰掉。也就是说，当缓存满的时候，应当把最先进入缓存的数据给淘汰掉。在FIFO Cache中应该支持以下操作;
Get(key)：如果Cache中存在该key，则返回对应的value值，否则，返回-1；
Put(key,value)：如果Cache中存在该key，则重置value值；如果不存在该key，则将该key插入到到Cache中，若Cache已满，则淘汰最早进入Cache的数据。
*/

type FIFOCache struct {
	Size     int //fifo的长度 , Size <= Capacity
	Capacity int //fifo的容量 ,初始化时，此值必须 >0
	Head     *DoubleNode
	Tail     *DoubleNode
	mutex    *sync.Mutex                 //读写并发控制
	Nodes    map[interface{}]*DoubleNode //map的key就是DoubleNode里的key，实现查找结点时间复杂度为 O(1)
}

func initFIFO(capacity int) (fifo *FIFOCache, err error) {
	if capacity <= 0 {
		return nil, errors.New("capacity <= 0")
	}
	return &FIFOCache{
		Size:     0,
		Capacity: capacity,
		Head:     nil,
		Tail:     nil,
		mutex:    new(sync.Mutex),
		Nodes:    make(map[interface{}]*DoubleNode, capacity),
	}, nil
}

func (fifo *FIFOCache) Put(key interface{}, value *DoubleNode) {
	fifo.mutex.Lock()
	defer func() {
		fifo.mutex.Unlock()
	}()

	if key == nil || value == nil {
		return
	}
	if fifo.Head == nil && fifo.Tail == nil {
		//第一次插入
		fifo.Size++
		fifo.Head = value
		fifo.Tail = value
		fifo.Nodes[key] = value
		return
	}

	if _, ok := fifo.Nodes[key]; ok {
		fifo.Nodes[key] = value
		return
	} else {
		if fifo.Size >= fifo.Capacity {
			fifo.removeLastNode()
		}
		fifo.putToHead(key, value)
		return
	}
}

//移除最后一个node
func (fifo *FIFOCache) removeLastNode() {
	if fifo.Tail == nil {
		return
	}
	delete(fifo.Nodes, fifo.Tail.Key)
	fifo.Size--
	pre := fifo.Tail.PrevNode
	fifo.Tail = pre
	fifo.Tail.PrevNode = pre.PrevNode
	fifo.Tail.NextNode = nil
}

//添加node到头部
func (fifo *FIFOCache) putToHead(key interface{}, value *DoubleNode) {
	if key == nil || value == nil {
		return
	}
	value.NextNode = fifo.Head
	value.PrevNode = nil
	fifo.Head.PrevNode = value
	fifo.Head = value
	fifo.Nodes[key] = value
	fifo.Size ++
}