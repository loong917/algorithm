package algorithm

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
)

type SkipList struct {
	head  *skipListNode // 头节点
	level int           // 当前层数
}

type skipListNode struct {
	data int             // 节点值
	next []*skipListNode // 当前节点的下一个节点（所有索引层）
}

const (
	MAX_LEVEL   = 6
	PROBABILITY = 0.25
)

func NewSkipList() *SkipList {
	//头结点
	head := newSkipListNode(math.MinInt, MAX_LEVEL)
	return &SkipList{
		head:  head,
		level: 1,
	}
}

func newSkipListNode(value int, level int) *skipListNode {
	return &skipListNode{
		data: value,
		next: make([]*skipListNode, level, level),
	}
}

func (list *SkipList) Find(value int) *skipListNode {
	node := list.head
	// 层数
	level := list.level
	// 查找路径
	paths := make([]string, 0)
	for i := level - 1; i >= 0; i-- {
		for node.next[i] != nil {
			paths = append(paths, fmt.Sprintf("%d（level%d）", node.next[i].data, i+1))
			if node.next[i].data == value {
				// 打印查找路径
				printPath(paths)
				return node.next[i]
			} else if node.next[i].data < value {
				node = node.next[i]
			} else if node.next[i].data > value {
				break
			}
		}
	}
	// 打印查找路径
	printPath(paths)
	return nil
}

func (list *SkipList) Insert(value int) {
	var level int = list.randomLevel()
	// new node
	newNode := newSkipListNode(value, level)
	node := list.head
	// store every level largest value which smaller than insert value
	floorNodes := make([]*skipListNode, MAX_LEVEL, MAX_LEVEL)
	for i := level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].data < value {
			node = node.next[i]
		}
		floorNodes[i] = node
	}
	for i := 0; i < level; i++ {
		newNode.next[i] = floorNodes[i].next[i]
		floorNodes[i].next[i] = newNode
	}
	// how to update level ?
	if list.level < level {
		list.level = level
	}
}

func (list *SkipList) Delete(value int) {
	node := list.head
	var level int = list.level
	// store every level largest value which smaller than insert value
	floorNodes := make([]*skipListNode, level, level)
	for i := level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].data < value {
			node = node.next[i]
		}
		floorNodes[i] = node
	}
	for i := level - 1; i >= 0; i-- {
		if floorNodes[i].next[i] != nil && floorNodes[i].next[i].data == value {
			floorNodes[i].next[i] = floorNodes[i].next[i].next[i]
			floorNodes[i] = nil
		}
	}
	// how to update level ?
	for list.level > 1 && list.head.next[list.level-1] == nil {
		list.level--
	}
}

// 随机算法获取该节点层数
func (list *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < PROBABILITY && level < MAX_LEVEL {
		level++
	}
	return level
}

func (list *SkipList) Print() string {
	if list.head == nil {
		return ""
	}
	result := make([]string, 0)
	// 数据
	var builder strings.Builder
	// 层数
	level := list.level
	for i := level - 1; i >= 0; i-- {
		builder.Reset()
		current := list.head
		for current.next[i] != nil {
			builder.WriteString(fmt.Sprintf("%+v", current.next[i].data))
			current = current.next[i]
			if current.next[i] != nil {
				builder.WriteString(" -> ")
			}
		}
		log.Printf("level%d = %+v \n", i+1, builder.String())
		result = append(result, builder.String())
	}
	return result[level-1]
}

// 打印查找路径
func printPath(nodes []string) {
	if len(nodes) == 0 {
		log.Println("NOT FOUND")
	}
	length := len(nodes)
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString(nodes[i])
		if i != length-1 {
			builder.WriteString(" -> ")
		}
	}
	log.Printf("查找路径 = %+v \n", builder.String())
}
