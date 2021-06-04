package structure

/**
队列
使用单链表结构
*/

type Queue struct {
	list *SingleList
}

//队列初始化
func (q *Queue) Init() {
	q.list = new(SingleList)
	q.list.Init()
}

//入队
//队尾插入
func (q *Queue) Enqueue(node *SingleNode) (bool, error) {
	return q.list.Insert(node, q.list.Size)
}

//出队
//队头取出
func (q *Queue) Dequeue() *SingleNode {
	node := q.list.Find(0)
	if node == nil {
		return nil
	}
	q.list.Delete(0)
	return node
}

//查看队头元素，并不删除
func (q *Queue) Peek() *SingleNode {
	node := q.list.Find(0)
	if node == nil {
		return nil
	}
	return node
}

//获取队列长度
func (q *Queue) Size() uint {
	return q.list.Size
}
