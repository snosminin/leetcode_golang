package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	previous1 := 1
	previous2 := 1

	for i := 2; i <= n; i++ {
		current := previous1 + previous2
		previous2 = previous1
		previous1 = current
	}

	return previous1
}
