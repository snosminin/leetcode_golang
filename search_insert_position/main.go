package main

import "fmt"

func main() {
	var array = [...]int{1, 3, 5, 6, 8}
	const numberToSearch = 2

	fmt.Println(BinarySearch(array[:], numberToSearch))
}

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := -1
	for left <= right {
		middle = (left + right) / 2

		if middle == left && middle == right {
			if target > nums[middle] {
				middle +=1
			}
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] == target {
			return middle
		}
	}

	return middle
}