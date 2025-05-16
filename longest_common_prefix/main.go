package main

import "fmt"

func main() {
	var array = []string{"flower","flow","flight"}

	fmt.Println(longestCommonPrefix(array))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	for symbolIndex := range 200 {
		symbolPrefixAccepted := true
		for stringIndex, str := range strs {
			if(stringIndex >= len(str) - 1) {
				return result
			}
			
			if(str[symbolIndex] != strs[stringIndex][symbolIndex]) {
				symbolPrefixAccepted = false
				break
			}
		}
		if(symbolPrefixAccepted) {
			result += string(strs[1][symbolIndex])
		}
	}
    return ""
}