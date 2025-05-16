package main

import "fmt"

func main() {
	var array = "MCMXCIV"

	fmt.Println(romanToInt(array))
}

func romanToInt(s string) int {
	result := 0
    var romanSymbols = map[string]int {
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
		"CM":900,
		"XC":90,
		"IX":9,
		"IV":4,
		"XL":40,
		"CD":400,
	}
	i := 0
	for i < len(s) {
		c := s[i]
		next := 0
		if(i < len(s) - 1) {
			val := romanSymbols[string(c) + string(s[i+1])]
			next = val
		}
		val, ok := romanSymbols[string(c)]
		if(next > 0) {
			result += next
			i += 1
		} else if(ok) {
			result += val
		}
		i += 1
	}

	return result
}