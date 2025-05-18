package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var l4 = ListNode{Val: 1, Next: nil}
	var l3 = ListNode{Val: 2, Next: &l4}
	var l2 = ListNode{Val: 2, Next: &l3}
	var l1 = ListNode{Val: 1, Next: &l2}

	fmt.Println(isPalindrome(&l1))
}

 
 func isPalindrome(head *ListNode) bool {
    return false
 }