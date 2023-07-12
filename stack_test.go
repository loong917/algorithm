package algorithm

import (
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	stack := NewArrayStack()
	stack.Push("杭州")
	stack.Push("北京")
	actual := stack.Print()
	if actual == expected {
		t.Logf("stack is '%s'", expected)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual, expected)
	}
}

func TestArrayStackPop(t *testing.T) {
	var (
		expected1 = "北京"
		expected2 = "武汉->杭州"
	)
	stack := NewArrayStack()
	stack.Push("武汉")
	stack.Push("杭州")
	stack.Push("北京")
	actual1, ok := stack.Pop()
	if !ok {
		t.Errorf("stack pop '%s'", "failed")
		return
	}
	if actual1 == expected1 {
		t.Logf("stack pop value is '%s'", expected1)
	} else {
		t.Errorf("stack pop value is '%s'; expected '%s'", actual1, expected1)
	}
	actual2 := stack.Print()
	if actual2 == expected2 {
		t.Logf("stack is '%s'", expected2)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual2, expected2)
	}
}

func TestLinkedListStackPush(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	stack := NewLinkedListStack()
	stack.Push("杭州")
	stack.Push("北京")
	actual := stack.Print()
	if actual == expected {
		t.Logf("stack is '%s'", expected)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual, expected)
	}
}

func TestLinkedListStackPop(t *testing.T) {
	var (
		expected1 = "北京"
		expected2 = "武汉->杭州"
	)
	stack := NewLinkedListStack()
	stack.Push("武汉")
	stack.Push("杭州")
	stack.Push("北京")
	actual1, ok := stack.Pop()
	if !ok {
		t.Errorf("stack pop '%s'", "failed")
		return
	}
	if actual1 == expected1 {
		t.Logf("stack pop value is '%s'", expected1)
	} else {
		t.Errorf("stack pop value is '%s'; expected '%s'", actual1, expected1)
	}
	actual2 := stack.Print()
	if actual2 == expected2 {
		t.Logf("stack is '%s'", expected2)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual2, expected2)
	}
}
