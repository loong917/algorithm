package algorithm

import (
	"errors"
	"fmt"
)

var ErrEmptyNode = errors.New(" The given node is null ")

var ErrNodeNotExist = errors.New(" The given node does not exist in list ")

var ErrRemoveFromEmptyList = errors.New(" It is not allowed to remove element from a empty list ")

// 单链表
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

// 在某个节点后面插入节点
func (list *SinglyLinkedList) InsertAfter(node *ListNode, v interface{}) error {
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
func (list *SinglyLinkedList) InsertBefore(node *ListNode, v interface{}) error {
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

// 打印链表结构
func (list *SinglyLinkedList) PrintAddress() string {
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
func (list *SinglyLinkedList) Reverse() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		return
	}
	fmt.Println("--- before reverse ---")
	fmt.Println(list.PrintAddress())
	var previous *ListNode
	var current *ListNode = list.head
	for current != nil {
		// 当前节点下一个节点
		temp := current.next
		current.next = previous
		// 更新上一个节点
		previous = current
		// 更新当前节点
		current = temp
	}
	list.head = previous
	fmt.Println("--- after reverse ---")
	fmt.Println(list.PrintAddress())
}

// 是否回文
func (list *SinglyLinkedList) IsPalindrome() bool {
	if list.head == nil {
		return false
	}
	if list.head.next == nil {
		return false
	}
	var isPalindrome bool = false
	fmt.Println("--- before verify palindrome ---")
	fmt.Println(list.PrintAddress())
	// 慢指针
	slow := list.head
	// 快指针
	fast := list.head
	// 当快指针遍历完链表，慢指针停下的位置就是中心点
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	// 反转链表后半部分
	var previous *ListNode
	for slow != nil {
		// 当前节点下一个节点
		tmp := slow.next
		slow.next = previous
		// 更新上一个节点
		previous = slow
		// 更新当前节点
		slow = tmp
	}
	// 链表左半部分
	left := list.head
	// 链表右半部分
	right := previous
	// 待还原链表
	restore := previous
	for right != nil {
		if left.data != right.data {
			isPalindrome = false
			goto RestoreLabel
		}
		left = left.next
		right = right.next
	}
	isPalindrome = true
RestoreLabel:
	// 还原链表
	var previousNode *ListNode
	for restore != nil {
		// 当前节点下一个节点
		tmp := restore.next
		restore.next = previousNode
		// 更新上一个节点
		previousNode = restore
		// 更新当前节点
		restore = tmp
	}
	fmt.Println("--- after verify palindrome ---")
	fmt.Println(list.PrintAddress())
	return isPalindrome
}

// 是否有环
func (list *SinglyLinkedList) HasRing() bool {
	if list.head == nil {
		return false
	}
	if list.head.next == nil {
		return false
	}
	// 慢指针
	slow := list.head
	// 快指针
	fast := list.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

