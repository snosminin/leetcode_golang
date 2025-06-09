package main

import "fmt"

func main() {
	const number = 58

	fmt.Println(intToRoman(number))
}

func intToRoman(num int) string {
	digits := num%10
	decades := num%100 - digits
	hundreds := num%1000 - decades - digits
	thousands := num - hundreds - decades - digits
	dictionary := map[int]string{
		4000 : "MMMM",
		3000 : "MMM",
		2000 : "MM",
		1000 : "M",
		900 : "CM",
		800 : "DCCC",
		700 : "DCC",
		600 : "DC",
		500 : "D",
		400 : "CD",
		300 : "CCC",
		200 : "CC",
		100 : "C",
		90 : "XC",
		80 : "LXXX",
		70 : "LXX",
		60 : "LX",
		50 : "L",
		40 : "XL",
		30 : "XXX",
		20 : "XX",
		10 : "X",
		9 : "IX",
		8 : "VIII",
		7 : "VII",
		6 : "VI",
		5 : "V",
		4 : "IV",
		3 : "III",
		2 : "II",
		1 : "I",
	}
	
	return dictionary[thousands]+dictionary[hundreds]+ dictionary[decades]+ dictionary[digits]
}