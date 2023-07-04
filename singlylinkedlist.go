package algorithm

import (
	"errors"
	"fmt"
)

var ErrEmptyNode = errors.New(" The given node is null ")

var ErrNodeNotExist = errors.New(" The given node does not exist in list ")

var ErrRemoveFromEmptyList = errors.New(" It is not allowed to remove element from a empty list ")

type SinglyLinkedList struct {
	head   *ListNode
	length uint
}

type ListNode struct {
	data interface{}
	next *ListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		length: 0,
	}
}

// 新增头节点
func (list *SinglyLinkedList) InsertToHead(v interface{}) error {
	// 前头节点
	next := list.head
	newNode := &ListNode{
		data: v,
		next: next,
	}
	list.head = newNode
	list.length++
	return nil
}

// 删除头节点
func (list *SinglyLinkedList) RemoveHead() error {
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	// 新头节点
	next := list.head.next
	list.head = next
	list.length--
	return nil
}

// 在某个节点后面插入节点
func (list *SinglyLinkedList) Insert(node *ListNode, v interface{}) error {
	if node == nil {
		return ErrEmptyNode
	}
	newNode := &ListNode{
		data: v,
		next: node.next,
	}
	node.next = newNode
	list.length++
	return nil
}

// 在某个节点前面插入节点
func (list *SinglyLinkedList) PreInsert(node *ListNode, v interface{}) error {
	// 空链表或者在头节点前插入
	if list.head == nil || node == list.head {
		return list.InsertToHead(v)
	}
	// 查找“前”节点
	var previous *ListNode
	current := list.head
	for current != node {
		previous = current
		current = current.next
	}
	if current == nil {
		return ErrNodeNotExist
	}
	newNode := &ListNode{
		data: v,
		next: current,
	}
	previous.next = newNode
	list.length++
	return nil
}

// 新增尾节点
func (list *SinglyLinkedList) InsertToTail(v interface{}) error {
	if list.head == nil {
		return list.InsertToHead(v)
	}
	// 查找“前”尾节点
	current := list.head
	for current.next != nil {
		current = current.next
	}
	newNode := &ListNode{
		data: v,
		next: nil,
	}
	current.next = newNode
	list.length++
	return nil
}

// 删除尾节点
func (list *SinglyLinkedList) RemoveTail() error {
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	// 查找“前”节点
	var previous *ListNode
	current := list.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	if previous == nil {
		// 有且仅有一个节点
		list.head = nil
	} else {
		// 不只一个节点
		previous.next = nil
	}
	list.length--
	return nil
}

// 删除节点
func (list *SinglyLinkedList) RemoveNode(node *ListNode) error {
	if node == nil {
		return ErrEmptyNode
	}
	if list.head == nil {
		return ErrRemoveFromEmptyList
	}
	previous := list.head
	current := list.head.next
	for current != nil {
		if current == node {
			break
		}
		previous = current
		current = current.next
	}
	if current == nil {
		return ErrNodeNotExist
	}
	previous.next = node.next
	list.length--
	return nil
}

// 打印链表
func (list *SinglyLinkedList) Print() string {
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
