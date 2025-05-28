package main

import "fmt"

func main() {
	var array = [...]int{10, 2, -2, -20, 10}
	const numberToSearch = -10

	fmt.Println(subarraySum(array[:], numberToSearch))
}

func subarraySum(nums []int, k int) int {
    prefixSum := 0
	result:=0
	prefixSumMap := make(map[int] int)

	for _, num := range nums {
		prefixSum += num
		if(prefixSum == k) {
			result++
		}
		count, ok := prefixSumMap[prefixSum - k] 
		if(ok) {
			result += count
		}

		countPrefixSum, okPrefixSum := prefixSumMap[prefixSum]
		if(!okPrefixSum){
			prefixSumMap[prefixSum] = 0
		}
		prefixSumMap[prefixSum] = countPrefixSum + 1
		

	}

	return result
}