package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	fmt.Println(inorderTraversal(nil))
}

func inorderTraversal(root *TreeNode) []int {
	if(root == nil) {
		return nil
	}

	result := []int {}
	if(root.Left != nil) {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if(root.Right != nil) {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}