package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	fmt.Println(isValidBST(nil))
}

func isValidBST(root *TreeNode) bool {
	return checkBST(root, nil, nil)
}

func checkBST(root *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if(root == nil) {
		return true
	}

	if(minNode!=nil && minNode.Val >= root.Val) {
		return false
	}
	if(maxNode!= nil && maxNode.Val <= root.Val) {
		return false
	}

	return checkBST(root.Left, minNode, root) && checkBST(root.Right, root, maxNode)
}