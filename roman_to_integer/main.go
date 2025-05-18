package main

import (
	"fmt"
)

func main() {
	var array = "MCMXCIV"

	fmt.Println(romanToInt(array))
}

type Pair struct {
    current byte
	next byte
}

func romanToInt(s string) int {
	result := 0
    var romanSymbols = map[byte]int {
		73:1,
		86:5,
		88:10,
		76:50,
		67:100,
		68:500,
		77:1000,
	}
		var romanSymbolPairs = map[Pair]int {
		{current: 67, next: 77}:900,
		{current: 88, next: 67}:90,
		{current: 73, next: 88}:9,
		{current: 73, next: 86}:4,
		{current: 88, next: 76}:40,
		{current: 67, next: 68}:400,
	}
	i := 0
	length := len(s)
	for i < length {
		c := s[i]

		if(i < length - 1){
			next := s[i+1]
			valPair, okPair := romanSymbolPairs[Pair{current: c, next: next}]
			if(okPair) {
				result += valPair
				i += 2
				continue
			}
		}

		val, ok := romanSymbols[c]
		if(ok) {
			result += val
		}
		i += 1
	}

	return result
}