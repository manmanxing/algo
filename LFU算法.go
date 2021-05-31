package main

import (
	"errors"
	"sync"
)

/**
LFU: 相对于 LRU，当数据容量已满，丢弃的数据是访问频次最低的数据
*/

//区别于普通的双向节点，LFU新增了一个频次参数
type LFUDoubleNode struct {
	*DoubleNode
	//访问频次
	Freq int
}

type LFUCache struct {
	Size     int //lfu的长度 , Size <= Capacity
	Capacity int //lfu的容量 ,初始化时，此值必须 >0
	Head     *LFUDoubleNode
	Tail     *LFUDoubleNode
	mutex    *sync.Mutex                    //读写并发控制
	Nodes    map[interface{}]*LFUDoubleNode //map的key 是 LFUDoubleNode 里的 Key
	FreMap   map[int]*DoubleList            //频次与对应的链表
}

func initLFU(capacity int) (lfu *LFUCache, err error) {
	if capacity <= 0 {
		return nil, errors.New("capacity <= 0")
	}

	return &LFUCache{
		Size:     0,
		Capacity: capacity,
		Head:     nil,
		Tail:     nil,
		mutex:    new(sync.Mutex),
		Nodes:    make(map[interface{}]*LFUDoubleNode, capacity),
		FreMap:   make(map[int]*DoubleList),
	}, nil
}

//更新某个节点的频次
func (lfu *LFUCache) updateFreq(node *LFUDoubleNode) {
	if node == nil {
		return
	}
	freq := node.Freq
	//找到对应频次的链表，将其在链表中删除
	lfu.FreMap[freq].Remove(node.DoubleNode)
	if lfu.FreMap[freq].Size == 0 {
		delete(lfu.FreMap, freq)
	}
	//增加频次，并插入该频次的链表中
	freq++
	node.Freq = freq

	if _, ok := lfu.FreMap[freq]; !ok {
		lfu.FreMap[freq], _ = InitDoubleList(10)
	}
	lfu.FreMap[freq].AddTail(node.DoubleNode)
}