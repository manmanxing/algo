package stackAndQueue

import "algo/list"

/*
	栈的实现
	使用单链表结构
*/

//栈信息
type Stack struct {
	list *list.SingleList
}

//栈的初始化
func (s *Stack) Init() {
	s.list = new(list.SingleList)
	s.list.Init()
}

//入栈
//这里将单链表的头结点看作栈顶
func (s *Stack) Push(node *list.SingleNode) (bool, error) {
	ok, err := s.list.Insert(node, 0)
	return ok, err
}

//出栈
func (s *Stack) Pop() *list.SingleNode {
	node := s.list.Find(0)
	if node == nil {
		return nil
	}
	s.list.Delete(0)
	return node
}

//查看栈顶结点
func (s *Stack) Peek() *list.SingleNode {
	node := s.list.Find(0)
	if node == nil {
		return nil
	}
	return node
}

//查看栈长度
func (s *Stack) Size() uint {
	return s.list.Size
}
