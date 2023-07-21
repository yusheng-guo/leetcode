package main

// 2427. 公因子的数目
// https://leetcode.cn/problems/number-of-common-factors/
func commonFactors(a int, b int) int {
	ans := 0
	a = gcd(a, b)
	for i := 1; i*i <= a; i++ {
		if a%i == 0 {
			ans++
			if a/i != i {
				ans++
			}
		}
	}
	return ans
}

// gcd 求最大公因数 方法二 更相减损法
//func gcd(a, b int) int {
//	for a != b {
//		if a > b {
//			a = a - b
//		} else {
//			b = b - a
//		}
//	}
//	return a
//}

// gcd 求最大公因数 方法一 辗转相除法 递归
//func gcd(a, b int) int {
//	if b == 0 {
//		return a
//	}
//	return gcd(b, a%b)
//}

// gcd 求最大公因数 方法一 辗转相除法 枚举
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// func main() {
// 	a, b := 927, 513
// 	ret := commonFactors(a, b)
// 	fmt.Println(ret)
// }
