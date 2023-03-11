package main

import "fmt"

// 219. 存在重复元素 II
// https://leetcode.cn/problems/contains-duplicate-ii/description/

// 方法三 滑动窗口
// 遍历nums
//
//	1.当i>k时 删除窗口左边元素 若nums[i]在窗口中，返回true；否则向窗口中添加元素nums[i]
//	2.当i<=k时 若nums[i]在窗口中，返回true；否则向窗口中添加元素nums[i]
func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]struct{}{}
	for i, n := range nums {
		if i > k {
			delete(window, nums[i-k-1])
		}
		if _, ok := window[n]; ok {
			return true
		}
		window[n] = struct{}{}
	}
	return false
}

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
