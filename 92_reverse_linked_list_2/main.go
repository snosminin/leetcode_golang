package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l5 = ListNode{Val: 5, Next: nil}
	var l4 = ListNode{Val: 4, Next: &l5}
	var l3 = ListNode{Val: 3, Next: &l4}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}

	fmt.Println(reverseBetween(&l1, 4, 5))
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	start := &ListNode{Val: 0, Next: head}
	current := start

	for i := 1; i < left; i++ {
		current = current.Next
	}

	beforeLeft := current
	current = current.Next

	stack := []*ListNode{}
	for i := left; i <= right; i++ {
		stack = append(stack, current)
		current = current.Next
	}

	startRight := current

	for i := left; i <= right; i++ {
		beforeLeft.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		beforeLeft = beforeLeft.Next
	}

	beforeLeft.Next = startRight

	return start.Next
}
