package main

import (
	"fmt"
)

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/
// 方法二 位运算(异或 满足交换律和结合律)
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}

// 方法一 map 哈希表
//func singleNumber(nums []int) int {
//	m := make(map[int]int, len(nums))
//	for _, v := range nums {
//		m[v]++
//	}
//	k := 0
//	for k = range m {
//		if m[k] == 1 {
//			break
//		}
//	}
//	return k
//}

func main() {
	nums := []int{2, 2, 1}
	ret := singleNumber(nums)
	fmt.Println(ret)
}
