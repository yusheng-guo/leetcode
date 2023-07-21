package main

// 2413. 最小偶倍数
// https://leetcode.cn/problems/smallest-even-multiple/
// 方法一
//
//	func smallestEvenMultiple(n int) int {
//		if n%2 == 0 {
//			return n
//		} else {
//			return n * 2
//		}
//	}
//
// 方法二 模运算
//func smallestEvenMultiple(n int) int {
//	return (n%2 + 1) * n
//}

// 方法三 与运算
//	func smallestEvenMultiple(n int) int {
//		return (n&1 + 1) * n
//	}
//

// 方法四 位运算
func smallestEvenMultiple(n int) int {
	return n << (n & 1)
}

// func main() {
// 	fmt.Println(smallestEvenMultiple(5))
// }
