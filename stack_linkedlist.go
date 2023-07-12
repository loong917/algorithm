package algorithm

import "fmt"

// 链式栈
type LinkedListStack struct {
	bottom *StackNode // 栈底
	top    *StackNode // 栈顶
}

type StackNode struct {
	data interface{}
	next *StackNode
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		bottom: nil,
		top:    nil,
	}
}

// 入栈
func (obj *LinkedListStack) Push(v interface{}) {
	newNode := &StackNode{
		data: v,
		next: nil,
	}
	// 空栈
	if obj.top == nil {
		obj.bottom = newNode
		obj.top = newNode
		return
	}
	obj.top.next = newNode
	obj.top = newNode
}

// 出栈
func (obj *LinkedListStack) Pop() (interface{}, bool) {
	if obj.top == nil {
		return nil, false
	}
	val := obj.top.data
	// 有且仅有一个节点
	if obj.bottom == obj.top {
		obj.bottom = nil
		obj.top = nil
		return val, true
	}
	// 查找“栈顶”前一个节点
	current := obj.bottom
	for current.next != obj.top {
		current = current.next
	}
	current.next = nil
	obj.top = current
	return val, true
}

// 打印栈
func (obj *LinkedListStack) Print() string {
	if obj.bottom == nil {
		return ""
	}
	current := obj.bottom
	var message string
	for current != nil {
		message += fmt.Sprintf("%+v", current.data)
		current = current.next
		if current != nil {
			message += "->"
		}
	}
	return message
}
