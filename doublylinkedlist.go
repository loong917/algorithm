package algorithm

import (
	"fmt"
)

// 双向链表
type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length uint
}

type Node struct {
	data     interface{}
	previous *Node
	next     *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		length: 0,
	}
}

// 新增头节点
func (list *DoublyLinkedList) InsertToHead(v interface{}) error {
	newNode := &Node{
		data: v,
	}
	// 头节点
	head := list.head
	if head == nil {
		// 空链表
		list.head = newNode
		list.tail = newNode
	} else {
		// 非空链表
		newNode.next = head
		head.previous = newNode
		list.head = newNode
	}
	list.length++
	return nil
}

// 删除头节点
func (list *DoublyLinkedList) RemoveHead() error {
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	// 新头节点
	next := list.head.next
	next.previous = nil
	list.head = next
	list.length--
	return nil
}

// 新增尾节点
func (list *DoublyLinkedList) InsertToTail(v interface{}) error {
	if list.head == nil {
		return list.InsertToHead(v)
	}
	tail := list.tail
	newNode := &Node{
		data:     v,
		previous: tail,
	}
	tail.next = newNode
	list.tail = newNode
	list.length++
	return nil
}

// 删除尾节点
func (list *DoublyLinkedList) RemoveTail() error {
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	tail := list.tail
	// 新的尾节点
	current := tail.previous
	current.next = nil
	list.tail = current
	list.length--
	return nil
}

// 在某个节点后面插入节点
func (list *DoublyLinkedList) InsertAfter(node *Node, v interface{}) error {
	if node == nil {
		return ErrEmptyNode
	}
	newNode := &Node{
		data: v,
	}
	tail := list.tail
	if node == tail {
		// 在尾节点后插入
		newNode.previous = tail
		tail.next = newNode
		list.tail = newNode
	} else {
		// 不是尾节点后插入
		newNode.previous = node
		newNode.next = node.next
		node.next = newNode
	}
	list.length++
	return nil
}

// 在某个节点前面插入节点
func (list *DoublyLinkedList) InsertBefore(node *Node, v interface{}) error {
	if node == nil {
		return ErrEmptyNode
	}
	// 空链表或者在头节点前插入
	if list.head == nil || node == list.head {
		return list.InsertToHead(v)
	}
	previous := node.previous
	newNode := &Node{
		data:     v,
		next:     node,
		previous: previous,
	}
	// 前一节点的next
	previous.next = newNode
	// 当前节点的previous
	node.previous = newNode
	list.length++
	return nil
}

// 删除节点
func (list *DoublyLinkedList) RemoveNode(node *Node) error {
	if node == nil {
		return ErrEmptyNode
	}
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	// 前一节点的next
	current := node.previous
	current.next = node.next
	// 后一节点的previous
	next := node.next
	next.previous = current
	list.length--
	return nil
}

// 打印链表
func (list *DoublyLinkedList) Print() string {
	if list.head == nil {
		return ""
	}
	current := list.head
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

// 打印链表结构
func (list *DoublyLinkedList) PrintAddress() string {
	if list.head == nil {
		return ""
	}
	current := list.head
	var address string
	for current != nil {
		address += fmt.Sprintf("%+v", current)
		current = current.next
		if current != nil {
			address += "->"
		}
	}
	return address
}

// 反转链表
func (list *DoublyLinkedList) Reverse() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		return
	}
	fmt.Println("--- before reverse ---")
	fmt.Println(list.PrintAddress())
	var previous *Node
	var current *Node = list.head
	for current != nil {
		// 当前节点下一个节点
		temp := current.next
		current.next = previous
		current.previous = temp
		if previous == nil {
			list.tail = current
		}
		// 更新上一个节点
		previous = current
		// 更新当前节点
		current = temp
	}
	list.head = previous
	fmt.Println("--- after reverse ---")
	fmt.Println(list.PrintAddress())
}
