package main

import (
	"fmt"
)

// 1018. 可被 5 整除的二进制前缀
// https://leetcode.cn/problems/binary-prefix-divisible-by-5/
// 优化
func prefixesDivBy5(nums []int) []bool {
	l := len(nums)
	ans := make([]bool, l)
	tmp := 0
	for k, v := range nums {
		tmp = (tmp<<1 | v) % 5
		ans[k] = tmp == 0
	}
	return ans
}

//func prefixesDivBy5(nums []int) []bool {
//	l := len(nums)
//	ans := make([]bool, l)
//	tmp := 0
//	for i := 0; i < l; i++ {
//		tmp = (tmp%5)<<1 | nums[i]
//		if tmp%5 == 0 {
//			ans[i] = true
//		} else {
//			ans[i] = false
//		}
//	}
//	return ans
//}

func main() {
	fmt.Println(prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0,
		1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1,
		0, 0, 1, 0}))
}
