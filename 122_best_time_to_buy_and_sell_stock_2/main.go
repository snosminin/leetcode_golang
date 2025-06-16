package main

import "fmt"

func main() {
	fmt.Println(maxProfit([] int{7,1,5,3,6,4}))
}

func maxProfit(prices []int) int {
	result := 0

	for i := 1; i < len(prices); i++ {
		dailyProfit := max(0, prices[i] - prices[i-1])
		result += dailyProfit
	}
	
	return result
}