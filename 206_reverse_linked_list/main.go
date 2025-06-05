package main

import "fmt"

type ListNode struct {
	Val int
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

	result := reverseList(&l1)

	fmt.Println(result)
}

func reverseList(head *ListNode) *ListNode {
    if(head == nil) {
		return head
	}
	stack := []ListNode {}
	current:=head
	for {
		stack = append(stack, *current)
		current = current.Next
		if(current == nil) {
			break
		}
	}

	newHead := &ListNode {Val: 0, Next: nil}
	newCurrent := newHead

	for {
		newCurrent.Next = &stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		newCurrent = newCurrent.Next
		newCurrent.Next = nil
		if(len(stack) == 0) {
			break
		}
	}

	return newHead.Next
}