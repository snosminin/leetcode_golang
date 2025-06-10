package main

import "fmt"

func main() {
	ransomNote, magazine := "aa" , "aab"

	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
    for _, c := range magazine {
		magazineMap[c]++
	}

	for _, c := range ransomNote {
		count, ok := magazineMap[c]
		if(!ok || count == 0) {
			return false
		}
		magazineMap[c]--
	}
	return true
}