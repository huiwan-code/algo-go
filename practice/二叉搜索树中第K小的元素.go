package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for {
			if root == nil {
				break
			}
			stack = append(stack, root)
			root = root.Left
		}
		lastestNode := stack[len(stack) - 1]
		stack = stack[:len(stack)-1]
		k -= 1
		if k == 0 {
			return lastestNode.Val
		}
		root = lastestNode.Right
	}
}
