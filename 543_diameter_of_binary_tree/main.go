package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	fmt.Println(diameterOfBinaryTree(nil))
}

var maxDiameter int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	depthSearch(root)
	return maxDiameter
}

func depthSearch(root *TreeNode) int {
	if(root == nil) {
		return 0
	}

	leftDepth := depthSearch(root.Left)
	rightDepth := depthSearch(root.Right)

	maxDiameter = max(maxDiameter, leftDepth + rightDepth)

    return max(leftDepth, rightDepth) + 1
}