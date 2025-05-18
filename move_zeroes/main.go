package main

import (
	"fmt"
)

func main() {
	var array = [...]int{0,1,0,3,12}
	moveZeroes(array[:])
	fmt.Println(array)
}

func moveZeroes(nums []int)  {
	if(len(nums) == 0) {
		return
	}
	
	for range len(nums) - 1 {
		for j := range len(nums) - 1 {
			if(nums[j] == 0) {
				swap := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = swap
			}
		}
	}
}