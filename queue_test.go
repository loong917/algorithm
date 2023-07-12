package algorithm

import (
	"testing"
)

func TestArrayQueueEnqueue(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	queue := NewArrayQueue()
	queue.Enqueue("杭州")
	queue.Enqueue("北京")
	actual := queue.Print()
	if actual == expected {
		t.Logf("queue is '%s'", expected)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual, expected)
	}
}

func TestArrayQueueDequeue(t *testing.T) {
	var (
		expected1 = "武汉"
		expected2 = "杭州->北京"
	)
	queue := NewArrayQueue()
	queue.Enqueue("武汉")
	queue.Enqueue("杭州")
	queue.Enqueue("北京")
	actual1, ok := queue.Dequeue()
	if !ok {
		t.Errorf("queue dequeue '%s'", "failed")
		return
	}
	if actual1 == expected1 {
		t.Logf("queue dequeue value is '%s'", expected1)
	} else {
		t.Errorf("queue dequeue value is '%s'; expected '%s'", actual1, expected1)
	}
	actual2 := queue.Print()
	if actual2 == expected2 {
		t.Logf("queue is '%s'", expected2)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual2, expected2)
	}
}

func TestLinkedListQueueEnqueue(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	queue := NewLinkedListQueue()
	queue.Enqueue("杭州")
	queue.Enqueue("北京")
	actual := queue.Print()
	if actual == expected {
		t.Logf("queue is '%s'", expected)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual, expected)
	}
}

func TestLinkedListQueueDequeue(t *testing.T) {
	var (
		expected1 = "武汉"
		expected2 = "杭州->北京"
	)
	queue := NewLinkedListQueue()
	queue.Enqueue("武汉")
	queue.Enqueue("杭州")
	queue.Enqueue("北京")
	actual1, ok := queue.Dequeue()
	if !ok {
		t.Errorf("queue dequeue '%s'", "failed")
		return
	}
	if actual1 == expected1 {
		t.Logf("queue dequeue value is '%s'", expected1)
	} else {
		t.Errorf("queue dequeue value is '%s'; expected '%s'", actual1, expected1)
	}
	actual2 := queue.Print()
	if actual2 == expected2 {
		t.Logf("queue is '%s'", expected2)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual2, expected2)
	}
}
