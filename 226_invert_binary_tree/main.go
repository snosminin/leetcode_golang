package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	fmt.Println(isSymmetric(nil))
}

func isSymmetric(root *TreeNode) bool {
    return checkSymmetry(root, root)
}

func checkSymmetry(node1 *TreeNode, node2 *TreeNode) bool {
	if(node1 == nil && node2 == nil) {
		return true
	}
	if(node1 == nil) != (node2 == nil) {
		return false
	}
	if(node1.Val != node2.Val){
		return false
	}

	return checkSymmetry(node1.Left, node2.Right) && checkSymmetry(node1.Right, node2.Left)
}