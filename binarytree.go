package algorithm

// 二叉树
type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

// 新增
func (tree *BinaryTree) Insert(data int) bool {
	// 新增节点
	node := &TreeNode{
		data: data,
	}
	if tree.root == nil {
		tree.root = node
		return true
	}
	current := tree.root
	for current != nil {
		if data == current.data {
			return false
		} else if data > current.data {
			// 没有右边子节点
			if current.right == nil {
				current.right = node
				break
			}
			current = current.right
		} else if data < current.data {
			// 没有左边子节点
			if current.left == nil {
				current.left = node
				break
			}
			current = current.left
		}
	}
	return true
}

// 删除
func (tree *BinaryTree) Delete(data int) {
	// 要删除节点的父节点
	var parent *TreeNode
	// 要删除的节点
	current := tree.root
	for current != nil && current.data != data {
		parent = current
		if data > current.data {
			current = current.right
		} else if data < current.data {
			current = current.left
		}
	}
	if current == nil {
		return
	}
	// 要删除的节点有2个子节点
	//
	// 首先查找删除节点的后继节点，即删除节点右子树中最小的节点；然后将删除节点后继节点的值覆盖删除节点，并将后继节点的右子树挂载到后继节点的父节点上；
	if current.left != nil && current.right != nil {
		// 查找右子树中最小的节点
		minParent := current
		minNode := current.right
		for minNode.left != nil {
			minParent = minNode
			minNode = minNode.left
		}
		// 将右子树中最小节点的值替换到当前节点中
		current.data = minNode.data
		current = minNode
		parent = minParent
	}
	// 要删除的节点是叶子节点或者只有1个子节点
	//
	// 情况1：如果删除节点是叶子节点，只需删除引用
	//
	// 情况2：如果删除节点只存在左子节点，将删除节点的左子树挂载到父节点上
	//
	// 情况3：如果删除节点只存在右子节点，将删除节点的右子树挂载到父节点上
	var child *TreeNode
	if current.left != nil {
		child = current.left
	} else if current.right != nil {
		child = current.right
	} else {
		child = nil
	}

	if parent == nil {
		tree.root = child
	} else if parent.left == current {
		parent.left = child
	} else if parent.right == current {
		parent.right = child
	}
}

// 前序遍历
//
// 对于树中的任意节点来说，先打印这个节点，然后再打印它的左子树，最后打印它的右子树
func PreOrderLoop(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	res := []int{}
	var loop func(*TreeNode)
	loop = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.data)
		if root.left != nil {
			loop(root.left)
		}
		if root.right != nil {
			loop(root.right)
		}
	}
	loop(node)
	return res
}

// 中序遍历
//
// 对于树中的任意节点来说，先打印它的左子树，然后再打印它本身，最后打印它的右子树
func InOrderLoop(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		return []int{node.data}
	}
	res := InOrderLoop(node.left)
	res = append(res, node.data)
	res = append(res, InOrderLoop(node.right)...)
	return res
}

// 后序遍历
//
// 对于树中的任意节点来说，先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身
func PostOrderLoop(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	res := []int{}
	var loop func(*TreeNode)
	loop = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.left != nil {
			loop(root.left)
		}
		if root.right != nil {
			loop(root.right)
		}
		res = append(res, root.data)
	}
	loop(node)
	return res
}
