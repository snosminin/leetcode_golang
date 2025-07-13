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
	var l1 = ListNode{Val: 0, Next: &l2}

	var m3 = ListNode{Val: 9, Next: nil}
	var m2 = ListNode{Val: 9, Next: &m3}
	var m1 = ListNode{Val: 9, Next: &m2}

	result := addTwoNumbers(&l1, &m1)

	fmt.Println(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0, Next: nil}

	current1 := l1
	current2 := l2
	currentResult := result
	overflow := 0
	for {
		val1 := 0
		val2 := 0

		if current1 != nil {
			val1 = current1.Val
			current1 = current1.Next
		}

		if current2 != nil {
			val2 = current2.Val
			current2 = current2.Next
		}

		sum := val1 + val2 + overflow
		overflow = sum / 10

		currentResult.Next = &ListNode{Val: sum % 10, Next: nil}
		currentResult = currentResult.Next

		if current1 == nil && current2 == nil {
			if overflow > 0 {
				currentResult.Next = &ListNode{Val: overflow, Next: nil}
			}
			break
		}

	}

	return result.Next
}
