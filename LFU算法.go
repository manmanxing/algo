package main

import (
	"errors"
	"sync"
)


//区别于普通的双向节点，LFU新增了一个频次参数
type LFUDoubleNode struct {
	*DoubleNode
	//访问频次
	Freq int
}

type LFUCache struct {
	Size     int                            //lfu的长度 , Size <= Capacity
	Capacity int                            //lfu的容量 ,初始化时，此值必须 >0
	mutex    *sync.Mutex                    //读写并发控制
	Nodes    map[interface{}]*LFUDoubleNode //map的key 是 LFUDoubleNode 里的 Key
	FreMap   map[int]*DoubleList            //频次与对应的链表
}

func InitLFUDoubleNode(key, value interface{}) *LFUDoubleNode {
	return &LFUDoubleNode{
		DoubleNode: InitDoubleNode(key, value),
		Freq:       0,
	}
}

func InitLFU(capacity int) (lfu *LFUCache, err error) {
	if capacity <= 0 {
		return nil, errors.New("capacity <= 0")
	}

	return &LFUCache{
		Size:     0,
		Capacity: capacity,
		mutex:    new(sync.Mutex),
		Nodes:    make(map[interface{}]*LFUDoubleNode, capacity),
		FreMap:   make(map[int]*DoubleList),
	}, nil
}

//更新某个已存在的节点的频次
//只能用于查询
func (lfu *LFUCache) updateFreqForGet(node *LFUDoubleNode) {
	if node == nil {
		return
	}

	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	freq := node.Freq
	//找到对应频次的链表，将其在链表中删除
	lfu.FreMap[freq].Remove(node.DoubleNode)
	if lfu.FreMap[freq].Size == 0 {
		delete(lfu.FreMap, freq)
	}
	//增加频次，并插入该频次的链表的头部，这样尾部是该频率中访问时间最旧的
	freq++
	node.Freq = freq

	if _, ok := lfu.FreMap[freq]; !ok {
		lfu.FreMap[freq], _ = InitDoubleList(10)
	}
	lfu.FreMap[freq].AddHead(node.DoubleNode)
}

//更新某个不存在的节点的频次
//只能用于新增
func (lfu *LFUCache) updateFreqForAdd(node *LFUDoubleNode) {
	if node == nil {
		return
	}
	node.Freq = 1
	if _, ok := lfu.FreMap[node.Freq]; !ok {
		lfu.FreMap[node.Freq], _ = InitDoubleList(10)
	}
	lfu.FreMap[node.Freq].AddHead(node.DoubleNode)
	lfu.Size++
}

func (lfu *LFUCache) removeFreqLowNode() {
	//找到频率最低的
	min := 0
	for k, v := range lfu.FreMap {
		_ = v
		if min > k {
			min = k
		}
	}

	lfu.FreMap[min].RemoveTail()
	if lfu.FreMap[min].Size == 0 {
		//删除这个 map 的key
		delete(lfu.FreMap, min)
	}
	lfu.Size--
}

//无论是 add 还是 get，对应的节点的频次都应该+1
//如果容量不足，删除频率最低，访问时间最旧的节点，再添加新节点
func (lfu *LFUCache) Add(key, value interface{}) {
	if key == nil && value == nil {
		return
	}

	if lfu.Capacity <= 0 {
		return
	}

	if v, ok := lfu.Nodes[key]; ok {
		//如果能找到
		v.value = value
		lfu.updateFreqForGet(v)
		return
	} else {
		//如果是新增
		//首先判断是否超容量，超了就删除频率最低，时间最旧的那个，再新增数据
		//没有超，就新增数据，并更新频次
		if lfu.Capacity <= lfu.Size {
			lfu.removeFreqLowNode()
		}
		newLFUNode := InitLFUDoubleNode(key, value)
		lfu.updateFreqForAdd(newLFUNode)
		return
	}
}

//无论是 add 还是 get，对应的节点的频次都应该+1
//如果能查询到，则返回对应的node
//如果不能查询到就返回 -1
func (lfu *LFUCache) Get(key interface{}) interface{} {
	if key == nil {
		return -1
	}

	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	if v,ok := lfu.Nodes[key];ok {
		lfu.updateFreqForGet(v)
	}

	return -1
}