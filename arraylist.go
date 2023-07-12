package algorithm

import (
	"fmt"
	"strings"
)

type ArrayList struct {
	data   []interface{}
	length int
}

// 增长因子
const (
	growthFactor = 2 // growth by 100%
)

func NewArrayList(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// 新增元素
func (list *ArrayList) Add(values ...interface{}) {
	fmt.Printf("before add slice: %v, length: %d, len: %d, cap: %d\n", list.data, list.length, len(list.data), cap(list.data))
	// 切片容量
	capacity := cap(list.data)
	var temp int = list.length + len(values)
	if temp >= capacity {
		// 达到容量后，根据增长因子进行扩容
		newCapacity := growthFactor * (capacity + len(values))
		newElements := make([]interface{}, list.length, newCapacity)
		copy(newElements, list.data[:list.length])
		list.data = newElements
	}
	for _, value := range values {
		list.data = append(list.data, value)
		list.length++
	}
	fmt.Printf("after add slice: %v, length: %d, len: %d, cap: %d\n", list.data, list.length, len(list.data), cap(list.data))
}

// 删除元素
func (list *ArrayList) Remove(index int) {
	fmt.Printf("before remove slice: %v, length: %d, len: %d, cap: %d\n", list.data, list.length, len(list.data), cap(list.data))
	if index < 0 || index >= list.length {
		return
	}
	// shift to the left by one (slow operation, need ways to optimize this)
	copy(list.data[index:], list.data[index+1:list.length])
	// erase last element (write zero value).
	list.data[list.length-1] = nil
	list.data = list.data[:list.length-1]
	list.length--
	fmt.Printf("after remove slice: %v, length: %d, len: %d, cap: %d\n", list.data, list.length, len(list.data), cap(list.data))
}

// 查找索引
func (list *ArrayList) IndexOf(value interface{}) int {
	if list.length == 0 {
		return -1
	}
	for index, element := range list.data {
		if element == value {
			return index
		}
	}
	return -1
}

// second return parameter is true if index is within bounds of the array and array is not empty,
// otherwise false.
func (list *ArrayList) Get(index int) (interface{}, bool) {
	if index < 0 || index >= list.length {
		return nil, false
	}
	return list.data[index], true
}

// 清空
func (list *ArrayList) Clear() {
	list.length = 0
	list.data = []interface{}{}
}

// 打印
func (list *ArrayList) Print() string {
	values := []string{}
	for _, value := range list.data {
		values = append(values, fmt.Sprintf("%v", value))
	}
	var message string = strings.Join(values, "->")
	return message
}
