package algorithm

import (
	"fmt"
	"strings"
)

// 循环队列
type CircleQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

// 循环队列当队列满时，tail指向的位置实际上没有存储数据的
// 循环队列会浪费一个单元存储空间
func NewCircleQueue(capacity int) *CircleQueue {
	payload := make([]interface{}, capacity)
	return &CircleQueue{
		data:     payload,
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// 入队
func (obj *CircleQueue) Enqueue(v interface{}) bool {
	fmt.Printf("before enqueue: %+v\n", obj)
	// 判断队列是否已满
	offset := (obj.tail + 1) % obj.capacity
	// 队列已满
	if offset == obj.head {
		return false
	}
	// 队列未满
	obj.data[obj.tail] = v
	obj.tail = offset
	fmt.Printf("after enqueue: %+v\n", obj)
	return true
}

// 出队
func (obj *CircleQueue) Dequeue() (interface{}, bool) {
	fmt.Printf("before dequeue: %+v\n", obj)
	// 队列为空
	if obj.head == obj.tail {
		return nil, false
	}
	v := obj.data[obj.head]
	// 重置
	obj.data[obj.head] = nil
	obj.head = (obj.head + 1) % obj.capacity
	fmt.Printf("after dequeue: %+v\n", obj)
	return v, true
}

// 队列大小
func (obj *CircleQueue) Size() int {
	if obj.tail > obj.head {
		return obj.tail - obj.head
	} else if obj.tail < obj.head {
		return obj.capacity + obj.tail - obj.head
	} else {
		return 0
	}
}

// 打印队列
func (obj *CircleQueue) Print() string {
	if obj.head == obj.tail {
		return ""
	}
	length := obj.Size()
	slices := make([]interface{}, length, length)
	for i := 0; i < obj.Size(); i++ {
		offset := (obj.head + i) % obj.capacity
		slices[i] = obj.data[offset]
	}
	values := []string{}
	for _, val := range slices {
		values = append(values, fmt.Sprintf("%v", val))
	}
	var message string = strings.Join(values, "->")
	return message
}
