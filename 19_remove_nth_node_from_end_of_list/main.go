package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l7 = ListNode{Val: 7, Next: nil}
	var l6 = ListNode{Val: 6, Next: &l7}
	var l5 = ListNode{Val: 5, Next: &l6}
	var l4 = ListNode{Val: 4, Next: &l5}
	var l3 = ListNode{Val: 3, Next: &l4}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}

	fmt.Println(removeNthFromEnd(&l1, 1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := GetLength(head)
	if length == n {
		return head.Next
	}

	current := head
	for range length - n - 1 {
		current = current.Next
	}

	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}

func GetLength(head *ListNode) int {
	if head == nil {
		return 0
	}

	current := head
	index := 0
	for {
		if current == nil {
			return index
		}
		index++
		current = current.Next
	}
}
