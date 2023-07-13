package algorithm

import (
	"fmt"
	"strings"
)

// 顺序队列
type ArrayQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	payload := make([]interface{}, capacity)
	return &ArrayQueue{
		data:     payload,
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// 入队
func (obj *ArrayQueue) Enqueue(v interface{}) bool {
	fmt.Printf("before enqueue: %+v\n", obj)
	if obj.tail == obj.capacity {
		// 队列已满
		if obj.head == 0 {
			return false
		}
		// 队列未满，数据迁移
		for i := obj.head; i < obj.tail; i++ {
			obj.data[i-obj.head] = obj.data[i]
		}
		obj.tail -= obj.head
		obj.head = 0
	}
	obj.data[obj.tail] = v
	obj.tail++
	fmt.Printf("after enqueue: %+v\n", obj)
	return true
}

// 出队
func (obj *ArrayQueue) Dequeue() (interface{}, bool) {
	fmt.Printf("before dequeue: %+v\n", obj)
	// 队列为空
	if obj.head == obj.tail {
		return nil, false
	}
	v := obj.data[obj.head]
	// 重置
	obj.data[obj.head] = nil
	obj.head++
	fmt.Printf("after dequeue: %+v\n", obj)
	return v, true
}

// 打印队列
func (obj *ArrayQueue) Print() string {
	if obj.head == obj.tail {
		return ""
	}
	values := []string{}
	for i := obj.head; i <= obj.tail-1; i++ {
		values = append(values, fmt.Sprintf("%+v", obj.data[i]))
	}
	var message string = strings.Join(values, "->")
	return message
}