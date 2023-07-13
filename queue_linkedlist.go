package algorithm

import "fmt"

// 链式队列
type LinkedListQueue struct {
	head *QueueNode // head指针，指向队头
	tail *QueueNode // tail指针，指向队尾
}

type QueueNode struct {
	data interface{}
	next *QueueNode
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
	}
}

// 入队
func (obj *LinkedListQueue) Enqueue(v interface{}) bool {
	newNode := &QueueNode{
		data: v,
		next: nil,
	}
	// 空队列
	if obj.tail == nil {
		obj.head = newNode
		obj.tail = newNode
		return true
	}
	// 非空队列
	obj.tail.next = newNode
	obj.tail = newNode
	return true
}

// 出队
func (obj *LinkedListQueue) Dequeue() (interface{}, bool) {
	// 队列为空
	if obj.head == nil {
		return nil, false
	}
	val := obj.head.data
	// 有且仅有一个节点
	if obj.head == obj.tail {
		obj.head = nil
		obj.tail = nil
		return val, true
	}
	obj.head = obj.head.next
	return val, true
}

// 打印队列
func (obj *LinkedListQueue) Print() string {
	if obj.head == nil {
		return ""
	}
	current := obj.head
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
