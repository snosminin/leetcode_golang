package main

import "fmt"

func main() {
	array := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(array))
}

func threeSum(nums []int) [][]int {
	pairsMap := make(map[int][]int)
	result := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pairsSum := nums[i] + nums[j]
			arr, ok := pairsMap[-pairsSum]
			if ok {
				for _, k := range arr {
					result = append(result, []int{nums[i], nums[j], nums[k]})

				}
			}
		}
		_, ok := pairsMap[nums[i]]
		if !ok {
			pairsMap[nums[i]] = []int{}
		}
		pairsMap[nums[i]] = append(pairsMap[nums[i]], i)
	}

	return result
}
