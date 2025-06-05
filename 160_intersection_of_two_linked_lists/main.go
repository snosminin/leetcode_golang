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
	
	var m3 = ListNode{Val: 3, Next: &l4}
	var m2 = ListNode{Val: 2, Next: &m3}
	var m1 = ListNode{Val: 1, Next: &m2}

	result := getIntersectionNode(&l1, &m1)

	fmt.Println(result)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if(headA == nil || headB == nil) {
		return nil
	}

	stackA := []*ListNode {}
	currentA:=headA
	for {
		if(currentA == nil) {
			break
		}
		stackA = append(stackA, currentA)
		currentA = currentA.Next
	}

	stackB := []*ListNode {}
	currentB:=headB
	for {
		if(currentB == nil) {
			break
		}
		stackB = append(stackB, currentB)
		currentB = currentB.Next
	}

	var leastLength int;
	if (len(stackA) > len(stackB)) {
		leastLength = len(stackB)
	} else {
		leastLength = len(stackA)
	}
	
	var result *ListNode = nil
	for range leastLength {
		lastStackA := stackA[len(stackA)-1]
		lastStackB := stackB[len(stackB)-1]

		if(lastStackA == lastStackB) {
			result = lastStackA
		} else {
			break
		}
		stackA = stackA[:len(stackA)-1]
		stackB = stackB[:len(stackB)-1]
	}

	return result
}