package main

import (
	"fmt"
)

// 1017. 负二进制转换
// https://leetcode.cn/problems/convert-to-base-2/description/
// 方法二 优化
func baseNeg2(n int) string {
	val := 0x55555555 ^ (0x55555555 - n)
	if val == 0 {
		return "0"
	}
	ret := ""
	for val > 0 {
		ret = string(rune((val&1)+'0')) + ret
		val >>= 1

	}
	return ret
}

// 方法一
//func baseNeg2(n int) string {
//	if n == 0 {
//		return "0"
//	}
//	ret := ""
//	for n != 0 {
//		absN := n
//		if n < 0 {
//			absN = -n
//		}
//		ret = string(rune(absN%2+'0')) + ret
//		n = int(math.Ceil(float64(n) / float64(-2)))
//	}
//	return ret
//}

func main() {
	fmt.Println(baseNeg2(4))
	fmt.Println(baseNeg2(3))
	fmt.Println(baseNeg2(2))
}
