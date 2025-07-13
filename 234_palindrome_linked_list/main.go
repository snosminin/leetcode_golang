package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l7 = ListNode{Val: 1, Next: nil}
	var l6 = ListNode{Val: 2, Next: &l7}
	var l5 = ListNode{Val: 2, Next: &l6}
	var l4 = ListNode{Val: 2, Next: &l5}
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

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func GetLength(head *ListNode) int {
	count := 1
	current := head
	for {
		if current.Next != nil {
			count++
			current = current.Next
			continue
		}
		return count
	}
}

func isPalindrome(head *ListNode) bool {
	length := GetLength(head)
	halfLength := length / 2
	currentNode := head
	isOddNumber := length%2 == 1

	if length == 1 {
		return true
	}

	stack := Stack{items: []int{}}
	for range halfLength {
		stack.Push(currentNode.Val)
		currentNode = currentNode.Next
	}

	if isOddNumber {
		currentNode = currentNode.Next
	}

	for range halfLength {
		if stack.Pop() != currentNode.Val {
			return false
		}
		currentNode = currentNode.Next
	}

	return true
}
