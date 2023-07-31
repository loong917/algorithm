package algorithm

import (
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	math.Sqrt(2)
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
