package algorithm

type Queue interface {
	// 入队
	Enqueue(v interface{}) bool
	// 出队
	Dequeue() (interface{}, bool)
}