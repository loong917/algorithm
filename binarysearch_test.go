package algorithm

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var (
		array    = []int{1, 2, 3, 4, 5, 6}
		expected = 4
	)
	var input int = 5
	actual := BinarySearch(array, input)
	if actual == expected {
		t.Logf("二分查找：the index of %d in %v is %d", input, array, actual)
	} else {
		t.Errorf("二分查找：the index of %d in %v is %d; expected %v", input, array, actual, expected)
	}
}
