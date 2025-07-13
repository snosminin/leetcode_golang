package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	return generateTriangle(numRows, [][]int{})
}

func generateTriangle(numRows int, triangle [][]int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	result := generateTriangle(numRows-1, triangle)

	resultRow := []int{1}
	for i := range numRows - 2 {
		resultRow = append(resultRow, result[numRows-2][i]+result[numRows-2][i+1])
	}
	resultRow = append(resultRow, 1)

	return append(result, resultRow)
}
