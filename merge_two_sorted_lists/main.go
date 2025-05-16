package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var l3 = ListNode{Val: 4, Next: nil}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}

	var m3 = ListNode{Val: 4, Next: nil}
	var m2 = ListNode{Val: 3, Next: &m3}
	var m1 = ListNode{Val: 1, Next: &m2}

	fmt.Println(mergeTwoLists(&l1, &m1))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var resultListNode *ListNode = &ListNode{Val: 0, Next: nil}
    var current  = resultListNode
	
	for list1 != nil && list2 != nil {
		if(list1.Val <= list2.Val) {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	
	if (list1 == nil) {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return resultListNode.Next
}