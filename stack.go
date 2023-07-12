package algorithm

type Stack interface {
	// 入栈
	Push(v interface{})
	// 出栈
	Pop() (interface{}, bool)
}
