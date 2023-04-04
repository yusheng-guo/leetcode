package main

import (
	"fmt"
)

// 7. 整数反转
// https://leetcode.cn/problems/reverse-integer/
// 方法二 优化
func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + (x % 10)
		x /= 10
	}
	if ans > (1<<31-1) || ans < (-1<<31) {
		return 0
	}
	return ans
}

// 方法一 整数反转
//func reverse(x int) int {
//	l := lenInt(x)
//	ans := 0
//	for i := l - 1; i >= 0; i-- {
//		ans += (x % 10) * pow10(i)
//		x /= 10
//	}
//	if ans > (1<<31-1) || ans < (-1<<31) {
//		return 0
//	}
//	return ans
//}
//
//func pow10(n int) int {
//	ans := 1
//	for n > 0 {
//		ans *= 10
//		n--
//	}
//	return ans
//}
//
//func lenInt(n int) int {
//	if n == 0 {
//		return 1
//	}
//	if n < 0 {
//		n = -n
//	}
//	ans := 0
//	for n > 0 {
//		ans++
//		n /= 10
//	}
//	return ans
//}

func main() {
	x := -321
	ret := reverse(x)
	fmt.Println(ret)
}
