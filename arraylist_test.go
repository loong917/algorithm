package algorithm

import (
	"testing"
)

func TestArrayListAdd(t *testing.T) {
	var (
		expected = "功->成->不->必->在->我"
	)
	s1 := []interface{}{"功", "成"}
	list := NewArrayList(s1...)
	s2 := []interface{}{"不"}
	list.Add(s2...)
	s3 := []interface{}{"必"}
	list.Add(s3...)
	s4 := []interface{}{"在", "我"}
	list.Add(s4...)
	actual := list.Print()
	if actual == expected {
		t.Logf("arrayList is '%s'", expected)
	} else {
		t.Errorf("arrayList is '%s'; expected '%s'", actual, expected)
	}
}

func TestArrayListRemove(t *testing.T) {
	var (
		expected = "功->成->必->在->我"
	)
	s := []interface{}{"功", "成", "不", "必", "在", "我"}
	list := NewArrayList(s...)
	list.Remove(2)
	actual := list.Print()
	if actual == expected {
		t.Logf("arrayList is '%s'", expected)
	} else {
		t.Errorf("arrayList is '%s'; expected '%s'", actual, expected)
	}
}
