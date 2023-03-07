package main

import "fmt"

func romanToInt(s string) int {
	relation := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	ret := 0
	for i := 0; i < l-1; i++ {
		if relation[s[i]] < relation[s[i+1]] {
			ret -= relation[s[i]]
		} else {
			ret += relation[s[i]]
		}
	}
	return ret + relation[s[len(s)-1]]
}

func main() {
	s := "MCMXCIV-"
	ret := romanToInt(s)
	fmt.Println(ret)
}
