package algorithm

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var (
		in       = []int{3, 5, 2, 4, 1, 6}
		expected = []int{1, 2, 3, 4, 5, 6}
	)
	input := make([]int, len(in))
	// 切片拷贝
	copy(input, in)
	actual := BubbleSort(input, len(input))
	if reflect.DeepEqual(actual, expected) {
		t.Logf("%v 冒泡排序 %v", in, expected)
	} else {
		t.Errorf("%v 冒泡排序 %v; expected %v", in, actual, expected)
	}
}

func TestInsertionSort(t *testing.T) {
	var (
		in       = []int{3, 5, 2, 4, 1, 6}
		expected = []int{1, 2, 3, 4, 5, 6}
	)
	input := make([]int, len(in))
	// 切片拷贝
	copy(input, in)
	actual := InsertionSort(input, len(input))
	if reflect.DeepEqual(actual, expected) {
		t.Logf("%v 插入排序 %v", in, expected)
	} else {
		t.Errorf("%v 插入排序 %v; expected %v", in, actual, expected)
	}
}

func TestSelectionSort(t *testing.T) {
	var (
		in       = []int{3, 5, 2, 4, 1, 6}
		expected = []int{1, 2, 3, 4, 5, 6}
	)
	input := make([]int, len(in))
	// 切片拷贝
	copy(input, in)
	actual := SelectionSort(input, len(input))
	if reflect.DeepEqual(actual, expected) {
		t.Logf("%v 选择排序 %v", in, expected)
	} else {
		t.Errorf("%v 选择排序 %v; expected %v", in, actual, expected)
	}
}

func TestMergeSort(t *testing.T) {
	var (
		in       = []int{3, 5, 2, 4, 1, 6}
		expected = []int{1, 2, 3, 4, 5, 6}
	)
	input := make([]int, len(in))
	// 切片拷贝
	copy(input, in)
	actual := MergeSort(input)
	if reflect.DeepEqual(actual, expected) {
		t.Logf("%v 归并排序 %v", in, expected)
	} else {
		t.Errorf("%v 归并排序 %v; expected %v", in, actual, expected)
	}
}

func TestQuickSort(t *testing.T) {
	var (
		in       = []int{6, 11, 3, 9, 8}
		expected = []int{3, 6, 8, 9, 11}
	)
	input := make([]int, len(in))
	// 切片拷贝
	copy(input, in)
	QuickSort(input)
	if reflect.DeepEqual(input, expected) {
		t.Logf("%v 快速排序 %v", in, expected)
	} else {
		t.Errorf("%v 快速排序 %v; expected %v", in, input, expected)
	}
}
