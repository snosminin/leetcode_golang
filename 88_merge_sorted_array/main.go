package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	nums1, m, nums2, n := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1Index := m - 1
	nums2Index := n - 1
	mergedIndex := m + n - 1

	for nums2Index >= 0 {
		if(nums1Index < 0 || nums2[nums2Index] >= nums1[nums1Index]){
			nums1[mergedIndex] = nums2[nums2Index]
			nums2Index--
		} else {
			nums1[mergedIndex] = nums1[nums1Index]
			nums1Index--
		}
		mergedIndex--
	}
}