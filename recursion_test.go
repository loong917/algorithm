package algorithm

import (
	"testing"
)

func TestFootStep(t *testing.T) {
	var (
		expected = 3
	)
	actual := FootStep(3)
	if actual == expected {
		t.Logf("走上 3 个 台阶有 %d 种走法", expected)
	} else {
		t.Errorf("走上 3 个 台阶有 %d 种走法; expected %d 种走法", actual, expected)
	}
}

func TestMoreFootStep(t *testing.T) {
	var (
		expected = 5
	)
	actual := FootStep(4)
	if actual == expected {
		t.Logf("走上 4 个 台阶有 %d 种走法", expected)
	} else {
		t.Errorf("走上 4 个 台阶有 %d 种走法; expected %d 种走法", actual, expected)
	}
}
