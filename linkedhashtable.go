package algorithm

import (
	"fmt"
	"strings"
)

type LinkedHashTable struct {
	table    []lruNode // 散列表
	head     *lruNode  // 头指针
	tail     *lruNode  // 尾指针
	capacity int       // 容量
	length   int       // 长度
}

type lruNode struct {
	key      int
	value    interface{}
	previous *lruNode // 双向链表前驱指针
	next     *lruNode // 双向链表后继指针
	hnext    *lruNode // 散列表拉链
}

// 默认容量
const (
	DEFAULT_LENGTH = 6
)

func NewLinkedHashTable(capacity int) *LinkedHashTable {
	return &LinkedHashTable{
		table:    make([]lruNode, DEFAULT_LENGTH),
		capacity: capacity,
		length:   0,
	}
}

func (list *LinkedHashTable) Get(key int) interface{} {
	if list.head == nil {
		return nil
	}
	if node := list.searchNode(key); node != nil {
		list.moveToTail(node)
		return node.value
	}
	return nil
}

func (list *LinkedHashTable) Put(key int, value interface{}) {
	if node := list.searchNode(key); node != nil {
		// 数据在 LRU 中，更新数据并移动到尾部
		node.value = value
		list.moveToTail(node)
		return
	}
	// 插入数据
	list.insertNode(key, value)
	if list.length > list.capacity {
		// 删除头部
		var headKey int = list.head.key
		list.removeNode(headKey)
	}
}

// 新增节点 (当前场景前提条件：节点不存在)
// 如果数据已经在缓存中，需要将其移动到双向链表的尾部；
// 如果数据不在缓存中，还要看缓存有没有满。如果满了，则将双向链表头部的结点删除，然后再将数据放到链表的尾部；如果没有满，就直接将数据放到链表的尾部。
func (list *LinkedHashTable) insertNode(key int, value interface{}) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}
	// 哈希值
	hashCode := hash(key)
	node := &list.table[hashCode]
	newNode.hnext = node.hnext
	list.table[hashCode].hnext = newNode
	// 新节点位于尾部
	tail := list.tail
	if tail == nil {
		// 空链表
		list.head = newNode
		list.tail = newNode
	} else {
		// 非空链表
		newNode.previous = tail
		tail.next = newNode
		list.tail = newNode
	}
	list.length++
}

// 删除节点
// 借助散列表，可以在 O(1) 时间复杂度里找到要删除的结点。
// 借助双向链表，可以通过前驱指针 O(1) 时间复杂度获取前驱结点，然后进行删除。
func (list *LinkedHashTable) removeNode(key int) {
	if list.head == nil {
		return
	}
	// 从哈希表中删除
	hashCode := hash(key)
	node := &list.table[hashCode]
	current := node.hnext
	for current != nil && current.key != key {
		node = current
		current = current.hnext
	}
	if current == nil {
		return
	}
	node.hnext = current.hnext
	// 从链表中删除
	// 前一节点的next
	node.previous.next = node.next
	// 后一节点的previous
	node.next.previous = node.previous
	// avoid memory leaks
	node.next = nil
	node.previous = nil
	node.value = nil
	list.length--
}

// 查找节点
// 散列表中查找数据的时间复杂度接近 O(1)，通过散列表可以很快地在缓存中找到一个数据。
// 当找到数据之后，需要将它移动到双向链表的尾部。
func (list *LinkedHashTable) searchNode(key int) *lruNode {
	if list.tail == nil {
		return nil
	}
	// 哈希值
	hashCode := hash(key)
	node := &list.table[hashCode]
	if node == nil {
		return nil
	}
	current := node.hnext
	for current != nil {
		if current.key == key {
			return current
		}
		current = current.hnext
	}
	return nil
}

// 移到尾部
func (list *LinkedHashTable) moveToTail(node *lruNode) {
	if list.tail == node {
		return
	}
	if list.head == node {
		list.head = node.next
		list.head.previous = nil
	} else {
		node.next.previous = node.previous
		node.previous.next = node.next
	}
	node.next = nil
	list.tail.next = node
	node.previous = list.tail
	// 尾节点
	list.tail = node
}

func (list *LinkedHashTable) Print() string {
	if list.head == nil {
		return ""
	}
	current := list.head
	// 值
	var message string
	for current != nil {
		message += fmt.Sprintf("%+v", current.value)
		current = current.next
		if current != nil {
			message += "->"
		}
	}
	return message
}

func (list *LinkedHashTable) PrintLinkedlist() string {
	if list.head == nil {
		return ""
	}
	current := list.head
	var builder strings.Builder
	for current != nil {
		builder.WriteString(fmt.Sprintf("%p【key：%d，value：%v，hash：%d，hnext：%p】", current, current.key, current.value, hash(current.key), current.hnext))
		current = current.next
		if current != nil {
			builder.WriteString("->")
		}
	}
	return builder.String()
}

func (list *LinkedHashTable) PrintHashtable() string {
	if list.head == nil {
		return ""
	}
	var address []string
	var builder strings.Builder
	for index, val := range list.table {
		builder.Reset()
		current := val.hnext
		for current != nil {
			builder.WriteString(fmt.Sprintf("%p", current))
			current = current.hnext
			if current != nil {
				builder.WriteString("->")
			}
		}
		if builder.Len() > 0 {
			address = append(address, fmt.Sprintf("%s【bucket%d】", builder.String(), index))
		}
	}
	return strings.Join(address, " => ")
}

// 哈希函数
func hash(key int) int {
	return (key ^ (key >> 32)) & (DEFAULT_LENGTH - 1)
}
