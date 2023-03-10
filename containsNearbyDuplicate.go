package main

import "fmt"

// 219. 存在重复元素 II
// https://leetcode.cn/problems/contains-duplicate-ii/description/

// 方法二 哈希表
//func containsNearbyDuplicate(nums []int, k int) bool {
//	m := map[int]int{}
//	for i := 1; i < len(nums)+1; i++ {
//		if m[nums[i-1]] != 0 && i-m[nums[i-1]] <= k {
//			return true
//		}
//		m[nums[i-1]] = i
//	}
//	return false
//}

// 方法一 哈希表+滑动窗口
// 问题：超时
//func containsNearbyDuplicate(nums []int, k int) bool {
//	l := len(nums)
//	if l <= k {
//		return contains(nums)
//	}
//	for i := 0; i < l-k; i++ {
//		if contains(nums[i : i+k+1]) {
//			return true
//		}
//	}
//	return false
//}
//
//func contains(nums []int) bool {
//	m := map[int]bool{}
//	for _, v := range nums {
//		if m[v] {
//			return true
//		} else {
//			m[v] = true
//		}
//	}
//	return false
//}

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	ret := containsNearbyDuplicate(nums, k)
	fmt.Println(ret)
}
