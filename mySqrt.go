package main

import (
	"fmt"
)

// x 的平方根
// https://leetcode.cn/problems/sqrtx/
// 方法四 牛顿迭代

// 方法三 二分查找 方法一优化
//func mySqrt(x int) int {
//	left, right := 0, x
//	mid := 0
//	for left <= right {
//		mid = (left + right) / 2
//		if mid*mid <= x && (mid+1)*(mid+1) > x {
//			return mid
//		}
//		if mid*mid < x {
//			left = mid + 1
//		}
//		if mid*mid > x {
//			right = mid - 1
//		}
//	}
//	return left
//}

// 方法二 其他数学函数代替
//func mySqrt(x int) int {
//	if x == 0 {
//		return 0
//	}
//	ans := int(math.Exp(0.5 * math.Log(float64(x))))
//	if (ans+1)*(ans+1) <= x {
//		return ans + 1
//	}
//	return ans
//}

// 方法一 暴力解法 时间复杂度 O(n)
//func mySqrt(x int) int {
//	for i := 0; i < x+1; i++ {
//		if i*i <= x && (i+1)*(i+1) > x {
//			return i
//		}
//	}
//	return -1
//}

func main() {
	fmt.Println(mySqrt(214323143))
}
