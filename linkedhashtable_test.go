package algorithm

import (
	"testing"
)

func TestLinkedHashTable(t *testing.T) {
	list := NewLinkedHashTable(10)
	list.Put(3, 11)
	list.Put(1, 12)
	list.Put(5, 23)
	list.Put(2, 22)
	t.Log("---before put---")
	t.Logf("%s", list.PrintHashtable())
	t.Logf("%s", list.PrintLinkedlist())
	list.Put(3, 26)
	t.Log("---after put---")
	t.Logf("%s", list.PrintHashtable())
	t.Logf("%s", list.PrintLinkedlist())
	var (
		expected = "12->23->22->26"
	)
	actual := list.Print()
	if actual == expected {
		t.Logf("linkedhashtable is '%s'", expected)
	} else {
		t.Errorf("linkedhashtable is '%s'; expected '%s'", actual, expected)
	}
	t.Log("---before get---")
	t.Logf("%s", list.PrintHashtable())
	t.Logf("%s", list.PrintLinkedlist())
	list.Get(5)
	t.Log("---after get---")
	t.Logf("%s", list.PrintHashtable())
	t.Logf("%s", list.PrintLinkedlist())
	var (
		expectedData = "12->22->26->23"
	)
	actualData := list.Print()
	if actualData == expectedData {
		t.Logf("linkedhashtable is '%s'", expectedData)
	} else {
		t.Errorf("linkedhashtable is '%s'; expected '%s'", actualData, expectedData)
	}

}
