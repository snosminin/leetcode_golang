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

type Stack struct {
    items []int
}

func (s *Stack) Push(data int) {
    s.items = append(s.items, data)
}

func (s *Stack) Pop() {
    if s.IsEmpty() {
        return
    }
    s.items = s.items[:len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

 
 func isPalindrome(head *ListNode) bool {
    return false
 }