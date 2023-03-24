package main

import (
	"fmt"
	"math"
)

// 414. 第三大的数
// https://leetcode.cn/problems/third-maximum-number/
// 方法三 方法二优化 一次循环
func thirdMax(nums []int) int {
	firstMaxNum, secondMaxNum, thirdMaxNum := math.MinInt, math.MinInt, math.MinInt
	for j := 0; j < len(nums); j++ {
		if nums[j] > firstMaxNum {
			thirdMaxNum = secondMaxNum
			secondMaxNum = firstMaxNum
			firstMaxNum = nums[j]
		}
		if nums[j] < firstMaxNum && nums[j] > secondMaxNum {
			thirdMaxNum = secondMaxNum
			secondMaxNum = nums[j]
		}
		if nums[j] < secondMaxNum && nums[j] > thirdMaxNum {
			thirdMaxNum = nums[j]
		}
	}
	if thirdMaxNum == math.MinInt {
		return firstMaxNum
	}
	return thirdMaxNum
}

// 方法二 三次循环
//func thirdMax(nums []int) int {
//	firstMaxNum, secondMaxNum, thirdMaxNum := math.MinInt, math.MinInt, math.MinInt
//	for j := 0; j < len(nums); j++ {
//		if nums[j] > firstMaxNum {
//			firstMaxNum = nums[j]
//		}
//	}
//	for j := 0; j < len(nums); j++ {
//		if nums[j] == firstMaxNum {
//			continue
//		}
//		if nums[j] > secondMaxNum {
//			secondMaxNum = nums[j]
//		}
//	}
//
//	for j := 0; j < len(nums); j++ {
//		if nums[j] == firstMaxNum || nums[j] == secondMaxNum {
//			continue
//		}
//		if nums[j] > thirdMaxNum {
//			thirdMaxNum = nums[j]
//		}
//	}
//	if thirdMaxNum == math.MinInt {
//		return firstMaxNum
//	}
//	return thirdMaxNum
//}

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
