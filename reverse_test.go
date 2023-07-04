package algorithm

import "testing"

func TestReverse(t *testing.T) {
	var (
		in       = "功成不必在我"
		expected = "我在必不成功"
	)
	actual := Reverse(in)
	if actual == expected {
		t.Logf("Passed：Reverse('%s') = '%s'", in, expected)
	} else {
		t.Errorf("Reverse('%s') = '%s'; expected = '%s'", in, actual, expected)
	}
}
