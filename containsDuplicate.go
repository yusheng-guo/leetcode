package main

import (
	"sort"
)

// 217. 存在重复元素
// https://leetcode.cn/problems/contains-duplicate/
// 方法二 排序
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 方法一 哈希表
//func containsDuplicate(nums []int) bool {
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

// func main() {
// 	nums := []int{1, 2, 3, 1}
// 	ret := containsDuplicate(nums)
// 	fmt.Println(ret)
// }
