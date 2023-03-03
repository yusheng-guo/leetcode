package main

import "fmt"

// 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/
// 方法二 二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		midv := nums[mid]
		if target == midv {
			return mid
		} else if target > midv {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

//1, 3, 6, 7
// 方法一 顺序查找
//func searchInsert(nums []int, target int) int {
//	l := len(nums)
//	for i := 0; i < l; i++ {
//		fmt.Println(i)
//		v := nums[i]
//		if v == target {
//			return i
//		}
//		if v < target {
//			continue
//		}
//		if v > target {
//			return i
//		}
//	}
//	return l
//}

func main() {
	nums := []int{1, 3, 6, 7}
	target := 6
	ret := searchInsert(nums, target)
	fmt.Println(ret)
}
