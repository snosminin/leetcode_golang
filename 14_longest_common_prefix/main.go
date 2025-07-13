package main

import "fmt"

func main() {
	var array = []string{"cir", "car"}

	fmt.Println(longestCommonPrefix(array))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for symbolIndex := range 200 {
		for stringIndex := range len(strs) - 1 {
			str := strs[stringIndex]
			strNext := strs[stringIndex+1]
			if symbolIndex >= len(str) || symbolIndex >= len(strNext) || str[symbolIndex] != strNext[symbolIndex] {
				return result
			}
		}
		result += string(strs[0][symbolIndex])
	}
	return result
}
