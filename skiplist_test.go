package algorithm

import (
	"log"
	"testing"
)

func TestSkipListFind(t *testing.T) {
	list := NewSkipList()
	for i := 1; i <= 20; i++ {
		list.Insert(i*2 + 1)
	}
	list.Insert(6)
	list.Insert(10)
	list.Insert(8)
	log.Printf("跳表：%+v\n", list.Print())
	var (
		in = 9
	)
	actual := list.Find(in)
	if actual != nil {
		t.Logf(" 跳表查找 %d 成功 ", in)
	} else {
		t.Errorf(" 跳表查找 %d 失败 ", in)
	}
}

func TestSkipListDelete(t *testing.T) {
	list := NewSkipList()
	for i := 1; i <= 20; i++ {
		list.Insert(i*2 + 1)
	}
	log.Printf("跳表：%+v\n", list.Print())
	var (
		in = 9
	)
	actual := list.Find(in)
	if actual != nil {
		log.Printf("跳表删除前：%+v\n", list.Print())
		log.Printf("跳表删除前：level = %d\n", list.level)
		list.Delete(in)
		log.Printf("跳表删除后：%+v\n", list.Print())
		log.Printf("跳表删除后：level = %d\n", list.level)
	} else {
		t.Errorf(" 跳表查找 %d 失败 ", in)
	}
}
