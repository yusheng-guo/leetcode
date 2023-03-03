package main

import (
	"fmt"
)

// 66. 加一
// https://leetcode.cn/problems/plus-one/
// 方法二
// 思路：从后☞前遍历数组 找9的个数
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 { // 获取9的个数
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 方法一
// 思路：将数组转换为数值 数值加一 将数值转化为数组
// 问题：数组长度有限
//func plusOne(digits []int) []int {
//	num := 0
//	l := len(digits)
//	for i := 0; i < l; i++ {
//		num = num*10 + digits[i]
//	}
//	num = num + 1
//	tmp := num
//	count := 0 // num位数
//	for tmp != 0 {
//		tmp /= 10
//		count++
//	}
//	if count > len(digits) {
//		digits = append(digits, num%10)
//		num /= 10
//	}
//	for i := 0; i < l; i++ {
//		digits[l-i-1] = num % 10
//		num /= 10
//	}
//	return digits
//}

func main() {
	nums := []int{9, 9, 9}
	ret := plusOne(nums)
	fmt.Println(ret)
}
