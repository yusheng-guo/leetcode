package main

import (
	"fmt"
)

// 剑指 Offer 56 - I. 数组中数字出现的次数
// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

// 方法三 分组异或 (思想：将数字分成两组 相同的数在同一个组 不同的数在不同的组)
func singleNumbers(nums []int) []int {
	// 全员异或
	ret := 0
	a, b := 0, 0 // 只出现了一次的两个数
	for _, v := range nums {
		ret ^= v
	}
	h := 1
	for ret&h == 0 {
		h = h << 1
	}
	for _, v := range nums {
		if v&h == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

// 方法二 排序
//func singleNumbers(nums []int) []int {
//	var ret []int
//	l := len(nums)
//	sort.Ints(nums)
//	if nums[0] != nums[1] {
//		ret = append(ret, nums[0])
//	}
//	if nums[l-1] != nums[l-2] {
//		ret = append(ret, nums[l-1])
//	}
//	for i := 1; i < len(nums)-1; i++ {
//		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
//			ret = append(ret, nums[i])
//		}
//	}
//	return ret
//}

// 方法一  哈希表
//func singleNumbers(nums []int) []int {
//	m := map[int]struct{}{}
//	for i := 0; i < len(nums); i++ {
//		if _, ok := m[nums[i]]; ok {
//			delete(m, nums[i])
//		} else {
//			m[nums[i]] = struct{}{}
//		}
//	}
//	var ret []int
//	for k := range m {
//		ret = append(ret, k)
//	}
//	return ret
//}

func main() {
	nums := []int{4, 1, 4, 6}
	ret := singleNumbers(nums)
	fmt.Println(ret)
}
