package main

import (
	"strconv"
)

// 1012. 至少有 1 位重复的数字
// https://leetcode.cn/problems/numbers-with-repeated-digits/
func numDupDigitsAtMostN(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if isContainsRepeatedNumbers(i) {
			count++
		}
	}
	return count
}

func isContainsRepeatedNumbers(n int) bool {
	if len(strconv.Itoa(100)) > 10 {
		return true
	}
	m := map[int]struct{}{}
	for n > 0 {
		if _, ok := m[n%10]; ok {
			return true
		}
		m[n%10] = struct{}{}
		n /= 10
	}
	return false
}

// func main() {
// 	fmt.Println(numDupDigitsAtMostN(100))
// }
