package algorithm

import (
	"testing"
)

func TestArrayQueueEnqueue(t *testing.T) {
	var (
		expected = "杭州->北京"
	)
	queue := NewArrayQueue(2)
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
	queue := NewArrayQueue(2)
	queue.Enqueue("武汉")
	queue.Enqueue("杭州")
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
	queue.Enqueue("北京")
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

func TestCircleQueueEnqueue(t *testing.T) {
	var (
		expected = "武汉->杭州"
	)
	queue := NewCircleQueue(3)
	queue.Enqueue("武汉")
	queue.Enqueue("杭州")
	// 队列已满
	ok := queue.Enqueue("北京")
	if !ok {
		t.Log("queue is full")
	} else {
		t.Errorf("queue enqueue '%s'; expected '%s'", "success", "failed")
	}
	actual := queue.Print()
	if actual == expected {
		t.Logf("queue is '%s'", expected)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual, expected)
	}
}

func TestCircleQueueDequeue(t *testing.T) {
	var (
		expected1 = "武汉"
		expected2 = "杭州->北京"
	)
	queue := NewCircleQueue(3)
	_, ok := queue.Dequeue()
	if !ok {
		t.Log("queue is empty")
	} else {
		t.Errorf("queue enqueue '%s'; expected '%s'", "success", "failed")
	}
	queue.Enqueue("武汉")            // head: 0, tail: 1
	queue.Enqueue("杭州")            // head: 0, tail: 2
	actual1, ok := queue.Dequeue() // head: 1, tail: 2
	if actual1 == expected1 {
		t.Logf("queue dequeue value is '%s'", expected1)
	} else {
		t.Errorf("queue dequeue value is '%s'; expected '%s'", actual1, expected1)
	}
	queue.Enqueue("北京") // head: 1, tail: 0
	actual2 := queue.Print()
	if actual2 == expected2 {
		t.Logf("queue is '%s'", expected2)
	} else {
		t.Errorf("queue is '%s'; expected '%s'", actual2, expected2)
	}
}