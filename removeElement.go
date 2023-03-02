package main

import "fmt"

// 27. 移除元素
// https://leetcode.cn/problems/remove-element/
// 方法二 左右双指针
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

// 方法一 快慢双指针
//func removeElement(nums []int, val int) int {
//	l := len(nums)
//	slow := 0
//	for fast := 0; fast < l; fast++ {
//		if nums[fast] != val {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}

func main() {
	nums := []int{2, 3, 4, 3, 3, 3, 4, 3}
	val := 3
	ret := removeElement(nums, val)
	fmt.Println(nums[:ret])
}
