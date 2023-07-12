package algorithm

import (
	"testing"
)

func TestArrayPush(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	stack := NewArrayStack()
	stack.Push("杭州")
	stack.Push("北京")
	actual := stack.Print()
	if actual == expected {
		t.Logf("arrayList is '%s'", expected)
	} else {
		t.Errorf("arrayList is '%s'; expected '%s'", actual, expected)
	}
}

func TestArrayPop(t *testing.T) {
	var (
		expected1 = "杭州->北京"
		expected2 = "北京"
		expected3 = "杭州"
	)
	stack := NewArrayStack()
	stack.Push("杭州")
	stack.Push("北京")
	actual1 := stack.Print()
	if actual1 == expected1 {
		t.Logf("stack is '%s'", expected1)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual1, expected1)
	}
	actual2, ok := stack.Pop()
	if !ok {
		t.Errorf("stack pop '%s'", "failed")
		return
	}
	if actual2 == expected2 {
		t.Logf("stack pop value is '%s'", expected2)
	} else {
		t.Errorf("stack pop value is '%s'; expected '%s'", actual2, expected2)
	}
	actual3 := stack.Print()
	if actual3 == expected3 {
		t.Logf("stack is '%s'", expected3)
	} else {
		t.Errorf("stack is '%s'; expected '%s'", actual3, expected3)
	}
}
