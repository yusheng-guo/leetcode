package main

// 461. 汉明距离
// https://leetcode.cn/problems/hamming-distance/
// 方法四 Brian Kernighan 算法
func hammingDistance(x int, y int) int {
	count := 0
	ret := x ^ y
	for ret != 0 {
		count++
		ret &= ret - 1
	}
	return count
}

// 方法三 异或 + 计数
//func hammingDistance(x int, y int) int {
//	count := 0
//	ret := x ^ y
//	for ret != 0 {
//		if ret&1 == 1 {
//			count++
//		}
//		ret >>= 1
//	}
//	return count
//}

// 方法二 内置函数
//func hammingDistance(x int, y int) int {
//	return bits.OnesCount(uint(x ^ y))
//}

// 方法一 移位
//func hammingDistance(x int, y int) int {
//	distance := 0
//	for x != 0 || y != 0 {
//		if x&1 != y&1 {
//			distance++
//		}
//		x >>= 1
//		y >>= 1
//	}
//	return distance
//}

// func main() {
// 	x := 1
// 	y := 4
// 	ret := hammingDistance(x, y)
// 	fmt.Println(ret)
// }
