package main

import "fmt"

func main() {
	haystack, needle := "mississippi", "issippi"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if(len(needle) > len(haystack)) {
		return -1
	}

    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		if(haystack[i] == needle[0]){
			if(haystack[i:i+len(needle)] == needle){
				return i
			}
		}
	}
	return -1
}