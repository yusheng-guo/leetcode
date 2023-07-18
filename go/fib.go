package main

import (
	"math"
)

// 方法三 通项公式
func fib(n int) int {
	sqrt5 := math.Sqrt(5)
	p1 := math.Pow((1+sqrt5)/2, float64(n))
	p2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((p1 - p2) / sqrt5))
}

// 方法二 迭代
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//	a, b := 0, 1
//	var ret int
//	for i := 1; i < n; i++ {
//		ret = b + a
//		b, a = ret, b
//	}
//	return ret
//}

// 方法一 递归
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//	return fib(n-1) + fib(n-2)
//}

// 1 1 2 3 5
// func main() {
// 	fmt.Println(fib(5))
// }
