package algorithm

import (
	"reflect"
	"testing"
)

func TestBinaryTreePreOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	// 前序遍历
	var root *TreeNode = tree.root
	actual := PreOrderLoop(root)
	var (
		expected = []int{4, 2, 1, 3, 6, 5, 7}
	)
	if reflect.DeepEqual(actual, expected) {
		t.Logf("前序遍历 %v", expected)
	} else {
		t.Errorf("前序遍历 %v; expected %v", actual, expected)
	}
}

func TestBinaryTreeInOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(16)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)
	tree.Insert(22)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(14)
	// 中序遍历
	var root *TreeNode = tree.root
	actual := InOrderLoop(root)
	var (
		expected = []int{6, 7, 8, 10, 12, 14, 16, 18, 20, 22}
	)
	if reflect.DeepEqual(actual, expected) {
		t.Logf("中序遍历 %v", expected)
	} else {
		t.Errorf("中序遍历 %v; expected %v", actual, expected)
	}
}

func TestBinaryTreePostOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	// 前序遍历
	var root *TreeNode = tree.root
	actual := PostOrderLoop(root)
	var (
		expected = []int{1, 3, 2, 5, 7, 6, 4}
	)
	if reflect.DeepEqual(actual, expected) {
		t.Logf("后序遍历 %v", expected)
	} else {
		t.Errorf("后序遍历 %v; expected %v", actual, expected)
	}
}

func TestBinaryTreeDelete(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	actual1 := InOrderLoop(tree.root)
	var (
		expected1 = []int{1, 2, 3, 4, 5, 6, 7}
	)
	if reflect.DeepEqual(actual1, expected1) {
		t.Logf("删除前中序遍历 %v", expected1)
	} else {
		t.Errorf("删除前中序遍历 %v; expected %v", actual1, expected1)
	}
	// 删除节点2
	tree.Delete(2)
	actual2 := InOrderLoop(tree.root)
	var (
		expected2 = []int{1, 3, 4, 5, 6, 7}
	)
	if reflect.DeepEqual(actual2, expected2) {
		t.Logf("删除后中序遍历 %v", expected2)
	} else {
		t.Errorf("删除后中序遍历 %v; expected %v", actual2, expected2)
	}
}
