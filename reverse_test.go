package algorithm

import "testing"

func TestReverseString(t *testing.T) {
	var (
		in       = "功成不必在我"
		expected = "我在必不成功"
	)
	actual := ReverseString(in)
	if actual == expected {
		t.Logf("reverse '%s' is '%s'", in, expected)
	} else {
		t.Errorf("reverse '%s' is '%s'; expected '%s'", in, actual, expected)
	}
}
