package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var o3 = ListNode{Val: 7, Next: nil}
	var o2 = ListNode{Val: 6, Next: &o3}
	var o1 = ListNode{Val: 2, Next: &o2}

	var l3 = ListNode{Val: 4, Next: nil}
	var l2 = ListNode{Val: 3, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}

	var m3 = ListNode{Val: 5, Next: nil}
	var m2 = ListNode{Val: 4, Next: &m3}
	var m1 = ListNode{Val: 1, Next: &m2}

	arr := []*ListNode {&l1, &m1, &o1}

	fmt.Println(mergeKLists(arr))
}

func mergeKLists(lists []*ListNode) *ListNode {
    if(lists == nil || len(lists) == 0) {
		return nil
	}

	resultListNode := &ListNode{Val: -11000, Next: nil}
	current := resultListNode

	for _, list := range lists {
		current = resultListNode
		for current != nil {
			if(list == nil) {
				break
			} else if(current.Next == nil) {
				current.Next = list
				list = list.Next
				current.Next.Next = nil
			} else if(list.Val >= current.Val && list.Val <= current.Next.Val) {
				swap := current.Next
				current.Next = list
				list = list.Next
				current.Next.Next = swap
			} else {
				current = current.Next 
			}
		}
	}

	return resultListNode.Next
}