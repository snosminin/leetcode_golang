package main

import "fmt"

func main() {
	var array = [...]int{1, 3, 5, 6, 8}
	const numberToSearch = 8

	fmt.Println(search(array[:], numberToSearch))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := -1
	for left <= right {
		middle = (left + right) / 2

		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] == target {
			return middle
		}
	}

	return -1
}
