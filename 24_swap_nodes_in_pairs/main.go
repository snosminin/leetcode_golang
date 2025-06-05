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

	fmt.Println(swapPairs(&l1))
}

func swapPairs(head *ListNode) *ListNode {
    if(head == nil) {
		return nil
	}
	if(head.Next == nil){
		return head
	}

	result := &ListNode {Val: 0, Next: head}
	current:=result

	for {
		if(current.Next == nil || current.Next.Next == nil) {
			return result.Next
		}

		nextNext:=current.Next.Next
		next:=current.Next
		swap := nextNext.Next
		current.Next = nextNext
		current.Next.Next = next
		current.Next.Next.Next = swap

		current = current.Next
		current = current.Next
	}
}