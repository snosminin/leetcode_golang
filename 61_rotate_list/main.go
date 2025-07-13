package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// var l8 = ListNode{Val: 8, Next: nil}
	// var l7 = ListNode{Val: 7, Next: &l8}
	// var l6 = ListNode{Val: 6, Next: &l7}
	// var l5 = ListNode{Val: 5, Next: &l6}
	// var l4 = ListNode{Val: 4, Next: &l5}
	// var l3 = ListNode{Val: 3, Next: &l4}
	var l2 = ListNode{Val: 2, Next: nil}
	var l1 = ListNode{Val: 1, Next: &l2}

	k := 1

	fmt.Println(rotateRight(&l1, k).Val)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	start := head
	current := start
	var end *ListNode = nil
	count := 0
	for {
		if current.Next == nil {
			end = current
			break
		}
		current = current.Next
		count++
	}
	count++

	k = k % count

	if k == 0 {
		return head
	}

	var result *ListNode = nil
	current = start
	cutIndex := count - k - 1
	i := 0
	for {
		if i == cutIndex {
			result = current.Next
			current.Next = nil
			end.Next = start
			break
		} else {
			current = current.Next
		}
		i++
	}

	return result
}
