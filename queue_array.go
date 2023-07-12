package algorithm

// 顺序队列
type ArrayQueue struct {
	list *ArrayList
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		list: NewArrayList(),
	}
}

// 入队
func (obj *ArrayQueue) Enqueue(v interface{}) {
	obj.list.Add(v)
}

// 出队
func (obj *ArrayQueue) Dequeue() (interface{}, bool) {
	value, ok := obj.list.Get(0)
	if ok {
		obj.list.Remove(0)
		return value, true
	}
	return value, false
}

// 打印队列
func (obj *ArrayQueue) Print() string {
	return obj.list.Print()
}
