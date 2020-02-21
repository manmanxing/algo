package main

import (
	"math/rand"
	"sync"
)

const (
	MAXLEVEL int     = 16   //可以存储 2^16 个元素
	P        float32 = 0.25 //p 是一个概率因子，它决定了随机生成 level 的大小分布。
)

/**
跳跃表
*/

type SkipData interface {
}

type SkipNode struct {
	Key       int         //排序所用字段，一般是下标索引
	Data      SkipData    //实际存储的数据，可以为任何类型
	NextNodes []*SkipNode //保存每层该结点指向下一个结点的指针
}

func newSkipNode(key, level int, data SkipData) *SkipNode {
	return &SkipNode{
		Key:       key,
		Data:      data,
		NextNodes: make([]*SkipNode, level, level),
	}
}

type SkipList struct {
	Head   *SkipNode     //头节点
	Tail   *SkipNode     //尾节点
	Length int           //最底层结点数量，也就是原链表的长度
	Level  int           //跳跃表的层数
	Mutex  *sync.RWMutex //读写互斥锁，用于保证协程安全
}

func newSkipList(level int) *SkipList {
	list := new(SkipList)
	if level <= 0 {
		level = MAXLEVEL
	}
	list.Level = level
	list.Head = newSkipNode(0, level, nil)
	list.Tail = newSkipNode(0, 0, nil)
	list.Mutex = new(sync.RWMutex)
	list.Length = 0
	//设置每一层的尾结点
	for i := range list.Head.NextNodes {
		list.Head.NextNodes[i] = list.Tail
	}
	return list
}

//随机数生成器，用于生成随机层数，随机生成的层数要满足P=0.5的几何分布
func randomLevel() int {
	level := 1
	for rand.Float32() < P && level < MAXLEVEL {
		level++
	}
	return level
}

//插入
func (s *SkipList) Insert(key int, data SkipData) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	//确定插入的深度
	level := randomLevel()
	//记录每层比key大或者是tail结点的上一个结点,索引代表层数
	preNodes := make([]*SkipNode, level, level)
	node := s.Head
	//查找插入的部位
	//从顶层开始往下遍历
	for i := level - 1; i >= 0; i-- {
		for {
			//每层头结点的下一个结点
			nextNode := node.NextNodes[i]
			if nextNode == s.Tail || nextNode.Key > key {
				//找到第一个比key大或者是tail结点
				//记录下该节点的前一个节点
				//然后执行下一层的比较
				preNodes[i] = node
				break
			} else if nextNode.Key == key {
				//如果刚好等于key，那么修改该结点的data
				//终止循环
				nextNode.Data = data
				return
			} else {
				//若node的下一个结点的集合中都没有符合的
				//那就移动到node的下一位结点A，再跟A的nextNode集合比较
				node = nextNode
			}
		}
	}
	//找到插入的部位后
	//将新的结点插入到每层符合条件的位置中
	//注意：这里的node是新结点插入的位置的前一个结点
	newNode := newSkipNode(key, level, data)
	for i, node := range preNodes {
		//交换位置
		node.NextNodes[i], newNode.NextNodes[i] = newNode, node.NextNodes[i]
	}
	s.Length++
}

//删除
func (s *SkipList) Delete(key int) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	//查找删除的结点
	//保存需要删除的结点的上一个结点,索引代表层数
	remove := make([]*SkipNode, s.Level, s.Level)
	node := s.Head
	var targetNode *SkipNode //用于保存删除的目标结点
	for i := len(node.NextNodes) - 1; i >= 0; i-- {
		for {
			nextNode := node.NextNodes[i]
			if nextNode == s.Tail || nextNode.Key > key {
				//说明该层没有找到目标,去下一层查询
				break
			} else if nextNode.Key == key {
				//如果刚好等于key，那么将上一个结点保存到集合，再去下一层查询
				remove[i] = node
				targetNode = nextNode
				break
			} else {
				//若node的下一个结点的集合中都没有符合的
				//那就移动到node的下一位结点A，再跟A的nextNode集合比较
				node = nextNode
			}
		}
	}
	//执行删除
	if len(remove) > 0 && targetNode != nil {
		for i, node := range remove {
			node.NextNodes[i] = targetNode.NextNodes[i]
		}
		s.Length--
		return true
	}
	return false
}

//查找
func (s *SkipList) Find(key int) SkipData {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	node := s.Head
	for i := len(node.NextNodes) - 1; i >= 0; i-- {
		for {
			nextNode := node.NextNodes[i]
			if nextNode == s.Tail || nextNode.Key > key {
				break
			} else if nextNode.Key == key {
				return nextNode.Data
			} else {
				node = nextNode
			}
		}
	}
	return nil
}

//获取数据总量
func (s *SkipList) Len() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return s.Length
}
