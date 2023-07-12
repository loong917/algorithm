package algorithm

// 顺序栈
type ArrayStack struct {
	list *ArrayList
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		list: NewArrayList(),
	}
}

// 入栈
func (obj *ArrayStack) Push(v interface{}) {
	obj.list.Add(v)
}

// 出栈
func (obj *ArrayStack) Pop() (interface{}, bool) {
	length := obj.list.length
	value, ok := obj.list.Get(length - 1)
	if ok {
		obj.list.Remove(length - 1)
		return value, true
	}
	return value, false
}

// 打印栈
func (obj *ArrayStack) Print() string {
	return obj.list.Print()
}
