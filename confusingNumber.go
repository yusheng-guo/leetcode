package main

import "fmt"

// 1056. 易混淆数
// https://leetcode.cn/problems/confusing-number/
func confusingNumber(n int) bool {
	m := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	origin := n  // 原始值
	reverse := 0 // 反转后的值
	for n > 0 {
		if _, ok := m[n%10]; !ok {
			return false
		}
		reverse = reverse*10 + m[n%10]
		n /= 10
	}
	return reverse != origin
}

func main() {
	n := 11
	ret := confusingNumber(n)
	fmt.Println(ret)
}
