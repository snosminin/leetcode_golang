package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var l8 = ListNode{Val: 1, Next: nil}
	var l7 = ListNode{Val: 1, Next: &l8}
	var l6 = ListNode{Val: 2, Next: &l7}
	var l5 = ListNode{Val: 2, Next: &l6}
	var l4 = ListNode{Val: 2, Next: &l5}
	var l3 = ListNode{Val: 2, Next: &l4}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}
	l8.Next = &l5

	fmt.Println(detectCycle(&l1))
}

func detectCycle(head *ListNode) *ListNode {
	if (head == nil) {
		return nil
	}
	current := head
	currentDoubleStep := head
    for current.Next != nil && currentDoubleStep.Next != nil {
		current = current.Next
		currentDoubleStep = currentDoubleStep.Next
		if(currentDoubleStep.Next == nil) {
			break
		}
		currentDoubleStep = currentDoubleStep.Next
		if(current == currentDoubleStep){
			return current
		}
	}
	return nil
}