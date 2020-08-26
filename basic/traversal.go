package algo

import "fmt"

type TreeNode struct {
	val    		int
	leftNode   	*TreeNode
	rightNode  	*TreeNode
}
// 前序排序
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preOrderTraversal(root.leftNode)
	preOrderTraversal(root.rightNode)
}

// 前序非递归
func preorderTraversal2(root *TreeNode) (ret []int) {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			ret = append(ret, root.val)
			stack = append(stack, root)
			root = root.leftNode
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		root = root.rightNode
	}
	return

}
// 中序排序
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.leftNode)
	fmt.Println(root.val)
	inorderTraversal(root.rightNode)
}

// 中序排序非递归
func inorderTraversal2(root *TreeNode) (ret []int){
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.leftNode
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret, root.val)
		root = root.rightNode
	}
	return
}


// 非递归后序
func postorderTraversal(root *TreeNode) (ret []int){
	if root == nil {
		return
	}
	var lastVisited *TreeNode
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.leftNode
		}
		root = stack[len(stack) - 1]
		if root.rightNode == nil ||  lastVisited == root.rightNode {
			stack = stack[:len(stack) - 1]
			ret = append(ret, root.val)
			lastVisited = root
		} else {
			root = root.rightNode
		}
	}
	return
}