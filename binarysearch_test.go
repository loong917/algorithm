package algorithm

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var (
		array    = []int{1, 3, 5, 7, 10, 14}
		expected = 4
	)
	var input int = 10
	actual := BinarySearch(array, input)
	if actual == expected {
		t.Logf("二分查找：the index of %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the index of %d in %v is %d; expected %v", input, array, actual, expected)
	}
}

func TestBinarySearchForRotateArray(t *testing.T) {
	var (
		array    = []int{5, 6, 7, 0, 1, 2, 3, 4}
		expected = 2
	)
	var input int = 7
	actual := BinarySearchForRotateArray(array, input)
	if actual == expected {
		t.Logf("二分查找：the index of %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the index of %d in %v is %d; expected %v", input, array, actual, expected)
	}
}

func TestBinarySearchFirst(t *testing.T) {
	var (
		array    = []int{1, 3, 5, 5, 7, 14}
		expected = 2
	)
	var input int = 5
	actual := BinarySearchFirst(array, input)
	if actual == expected {
		t.Logf("二分查找：the first index of %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the first index of %d in %v is %d; expected %v", input, array, actual, expected)
	}
}

func TestBinarySearchLast(t *testing.T) {
	var (
		array    = []int{1, 3, 5, 5, 7, 14}
		expected = 3
	)
	var input int = 5
	actual := BinarySearchLast(array, input)
	if actual == expected {
		t.Logf("二分查找：the last index of %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the last index of %d in %v is %d; expected %v", input, array, actual, expected)
	}
}

func TestBinarySearchFirstGE(t *testing.T) {
	var (
		array    = []int{1, 3, 5, 5, 7, 14}
		expected = 4
	)
	var input int = 6
	actual := BinarySearchFirstGE(array, input)
	if actual == expected {
		t.Logf("二分查找：the index of first element greater than or equal to %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the index of first element greater than or equal to %d in %v is %d; expected %v", input, array, actual, expected)
	}
}

func TestBinarySearchLastLE(t *testing.T) {
	var (
		array    = []int{1, 3, 5, 5, 7, 14}
		expected = 3
	)
	var input int = 6
	actual := BinarySearchLastLE(array, input)
	if actual == expected {
		t.Logf("二分查找：the index of last element less than or equal to %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the index of last element less than or equal to %d in %v is %d; expected %v", input, array, actual, expected)
	}
}
