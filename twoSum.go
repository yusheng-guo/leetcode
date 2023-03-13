package main

import (
	"fmt"
)

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/

// 方法三 哈希表
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, x := range nums {
		if _, ok := m[target-x]; ok {
			return []int{i, m[target-x]}
		}
		m[x] = i
	}
	return nil
}

// 方法二 双指针
// 时间复杂度 O(n*logn)
//func twoSum(nums []int, target int) []int {
//	l := len(nums)
//	var nums_ []int
//	for i := 0; i < l; i++ {
//		nums_ = append(nums_, nums[i])
//	}
//	sort.Ints(nums) // 排序
//	left, right := 0, len(nums)-1
//	for left < right {
//		if nums[left]+nums[right] == target {
//			break
//		}
//		if nums[left]+nums[right] < target {
//			left++
//		}
//		if nums[left]+nums[right] > target {
//			right--
//		}
//	}
//	if nums[left]+nums[right] != target {
//		return nil
//	}
//	p1, p2 := -1, -1
//	for i := 0; i < l; i++ {
//		if p1 == -1 && nums_[i] == nums[left] {
//			p1 = i
//			continue
//		}
//		if p2 == -1 && nums_[i] == nums[right] {
//			p2 = i
//		}
//	}
//	return []int{p1, p2}
//}

// 方法一 双重循环
// 时间复杂度 O(n*n)
//func twoSum(nums []int, target int) []int {
//	l := len(nums)
//	for i := 0; i < l; i++ {
//		for j := i + 1; j < l; j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 10
	ret := twoSum(nums, target)
	fmt.Println(ret)
}
