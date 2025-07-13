package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l8 = ListNode{Val: 8, Next: nil}
	var l7 = ListNode{Val: 7, Next: &l8}
	var l6 = ListNode{Val: 6, Next: &l7}
	var l5 = ListNode{Val: 5, Next: &l6}
	var l4 = ListNode{Val: 4, Next: &l5}
	var l3 = ListNode{Val: 3, Next: &l4}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}
	l8.Next = &l5

	fmt.Println(detectCycle(&l1))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	currentDoubleStep := head
	for current.Next != nil && currentDoubleStep.Next != nil {
		current = current.Next
		currentDoubleStep = currentDoubleStep.Next
		if currentDoubleStep.Next == nil {
			break
		}
		currentDoubleStep = currentDoubleStep.Next
		if current == currentDoubleStep {
			currentDoubleStep = head
			for current != currentDoubleStep {
				current = current.Next
				currentDoubleStep = currentDoubleStep.Next
			}
			return current
		}
	}
	return nil
}
