package main

import "fmt"

var ARRAY = [...]int{2, 7, 11, 15, 34, 2, 1, 6, 8, 4, 234, 2, 14, 7, 0}
const SUM_TO_SEARCH = 25

func main() {
	fmt.Println(twoSumSearch(ARRAY[:], SUM_TO_SEARCH))
}

func twoSumSearch(nums []int, target int) []int {
	var resultMap = make(map[int]int)
	for i, item := range nums {
		var diff = target - item
		_, ok := resultMap[diff]
		if ok {
			return []int {i, resultMap[diff]}
		}
		resultMap[item] = i
	}
	return nil
}