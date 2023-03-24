package main

import (
	"fmt"
	"math"
)

// 414. 第三大的数
// https://leetcode.cn/problems/third-maximum-number/
func thirdMax(nums []int) int {
	max(nums)
}

var count = 0
var thirdMaxNum = 0

func max(nums []int) (int, int) {
	m := math.MinInt
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
			index = i
		}
	}
	if index == len(nums)-1 {
		nums = nums[:index]
	}
	nums = append(nums[:index], nums[index+1:]...)
	count++
	if count > 2 {
		thirdMaxNum = m
	}
	return max(nums)
}

// 方法一：排序+查找
//func thirdMax(nums []int) int {
//	// 排序 从大到小
//	sort.Slice(nums, func(i, j int) bool {
//		if nums[i] > nums[j] {
//			return true
//		}
//		return false
//	})
//	count := 1
//	curr := nums[0]
//	fmt.Println(nums)
//	for i := 1; i < len(nums); i++ {
//		if curr > nums[i] {
//			count++
//			curr = nums[i]
//		}
//		if count > 2 {
//			return curr
//		}
//	}
//	return nums[0]
//}

func main() {
	arr := []int{3, 2, 1}
	ret := thirdMax(arr)
	fmt.Println(ret)
}
