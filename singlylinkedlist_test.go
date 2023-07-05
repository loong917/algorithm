package algorithm

import (
	"errors"
	"testing"
)

func TestEmptySinglyInsertToHead(t *testing.T) {
	list := NewSinglyLinkedList()
	var (
		in       = "1"
		expected = "1"
	)
	err := list.InsertToHead(in)
	if err != nil {
		t.Errorf("insert to head：%s", err.Error())
		return
	}
	actual := list.Print()
	if actual == expected {
		t.Logf("insert '%s' to head is '%s'", in, expected)
	} else {
		t.Errorf("insert '%s' to head is '%s'; expected '%s'", in, actual, expected)
	}
}

func TestSinglyInsertToHead(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToHead("1")
	if err != nil {
		t.Errorf("insert to head：%s", err.Error())
		return
	}
	t.Logf("before insert to head is '%s' ", list.Print())
	err = list.InsertToHead("2")
	if err != nil {
		t.Errorf("insert to head： %s", err.Error())
		return
	}
	var expected string = "2->1"
	actual := list.Print()
	if actual == expected {
		t.Logf("after insert to head is '%s' ", expected)
	} else {
		t.Errorf("insert to head is '%s'; expected '%s'", actual, expected)
	}
}

func TestEmptySinglyRemoveHead(t *testing.T) {
	list := NewSinglyLinkedList()
	var expected error = ErrRemoveFromEmptyList
	err := list.RemoveHead()
	if err == nil {
		t.Errorf("remove head is %v'; expected %v", err, expected)
		return
	}
	if errors.Is(err, expected) {
		t.Logf("remove head：%v", err)
	} else {
		t.Errorf("remove head：%v", err)
	}
}

func TestSinglyRemoveHead(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToHead("1")
	if err != nil {
		t.Errorf("remove head：%s", err.Error())
		return
	}
	err = list.InsertToHead("2")
	if err != nil {
		t.Errorf("remove head：%s", err.Error())
		return
	}
	t.Logf("before remove head is '%s'", list.Print())
	err = list.RemoveHead()
	if err != nil {
		t.Errorf("remove head：%s", err.Error())
		return
	}
	var expected string = "1"
	actual := list.Print()
	if actual == expected {
		t.Logf("after remove head is '%s'", expected)
	} else {
		t.Errorf("remove head is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyInsertToTail(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("insert to tail：%s", err.Error())
		return
	}
	t.Logf("before insert to tail is '%s'", list.Print())
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("insert to tail：%s", err.Error())
		return
	}
	var expected string = "1->2"
	actual := list.Print()
	if actual == expected {
		t.Logf("after insert to tail is '%s'", expected)
	} else {
		t.Errorf("insert to tail is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyRemoveTail(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("remove tail：%s", err.Error())
		return
	}
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("remove tail：%s", err.Error())
		return
	}
	t.Logf("before remove tail is '%s'", list.Print())
	err = list.RemoveTail()
	if err != nil {
		t.Errorf("remove tail：%s", err.Error())
		return
	}
	var expected string = "1"
	actual := list.Print()
	if actual == expected {
		t.Logf("after remove tail is '%s'", expected)
	} else {
		t.Errorf("remove tail is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyRemoveNode(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("remove node：%s", err.Error())
		return
	}
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("remove node：%s", err.Error())
		return
	}
	err = list.InsertToTail("4")
	if err != nil {
		t.Errorf("remove node：%s", err.Error())
		return
	}
	t.Logf("before remove node is '%s'", list.Print())
	var (
		in       = "2"
		expected = "1->4"
	)
	current := list.head
	var length uint = list.length
	var i uint = 0
	var node *ListNode
	for ; i < length; i++ {
		if current.data == in {
			node = current
			break
		}
		current = current.next
	}
	err = list.RemoveNode(node)
	if err != nil {
		t.Errorf("remove node：%s", err.Error())
		return
	}
	actual := list.Print()
	if actual == expected {
		t.Logf("after remove node is '%s'", expected)
	} else {
		t.Errorf("remove node is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyInsertAfter(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	err = list.InsertToTail("3")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	t.Logf("before insert node is '%s'", list.Print())
	var (
		in       = "3"
		expected = "1->2->3->4"
	)
	current := list.head
	var length uint = list.length
	var i uint = 0
	var node *ListNode
	for ; i < length; i++ {
		if current.data == in {
			node = current
			break
		}
		current = current.next
	}
	err = list.InsertAfter(node, "4")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	actual := list.Print()
	if actual == expected {
		t.Logf("after insert node is '%s'", expected)
	} else {
		t.Errorf("insert node is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyInsertBefore(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	err = list.InsertToTail("3")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	t.Logf("before insert node is '%s'", list.Print())
	var (
		in       = "3"
		expected = "1->2->4->3"
	)
	current := list.head
	var length uint = list.length
	var i uint = 0
	var node *ListNode
	for ; i < length; i++ {
		if current.data == in {
			node = current
			break
		}
		current = current.next
	}
	err = list.InsertBefore(node, "4")
	if err != nil {
		t.Errorf("insert node：%s", err.Error())
		return
	}
	actual := list.Print()
	if actual == expected {
		t.Logf("after insert node is '%s'", expected)
	} else {
		t.Errorf("insert node is '%s'; expected '%s'", actual, expected)
	}
}

func TestSinglyReverse(t *testing.T) {
	list := NewSinglyLinkedList()
	var err error
	err = list.InsertToTail("1")
	if err != nil {
		t.Errorf("reverse：%s", err.Error())
		return
	}
	err = list.InsertToTail("2")
	if err != nil {
		t.Errorf("reverse：%s", err.Error())
		return
	}
	err = list.InsertToTail("3")
	if err != nil {
		t.Errorf("reverse：%s", err.Error())
		return
	}
	t.Logf("before reverse is '%s'", list.Print())
	var expected string = "3->2->1"
	list.Reverse()
	actual := list.Print()
	if actual == expected {
		t.Logf("after reverse is '%s'", expected)
	} else {
		t.Errorf("reverse is '%s'; expected '%s'", actual, expected)
	}
}
