package main

import "fmt"

// 263. 丑数
// https://leetcode.cn/problems/ugly-number/description/
// 方法二 n=2**a × 3**b × 5**c
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	factors := []int{2, 3, 5}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

// 方法一 获取所有质因数进行比对
//func primeFactors(n int) []int {
//	factors := make([]int, 0)
//	for n%2 == 0 {
//		factors = append(factors, 2)
//		n = n / 2
//	}
//	for i := 3; i*i <= n; i = i + 2 {
//		for n%i == 0 {
//			factors = append(factors, i)
//			n = n / i
//		}
//	}
//	if n > 2 {
//		factors = append(factors, n)
//	}
//	return factors
//}
//
//func isUgly(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	factors := primeFactors(n)
//	fmt.Println(factors)
//	for _, v := range factors {
//		if v != 2 && v != 3 && v != 5 {
//			return false
//		}
//	}
//	return true
//}

func main() {
	n := 1
	fmt.Println(isUgly(n))
}
