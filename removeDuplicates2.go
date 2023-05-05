package main

import "fmt"

// 删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates2(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	// slow 已整理完成的数组长度
	// slow-1 被保留元素位置
	// slow-2 上次被保留元素位置
	slow := 2
	for fast := 2; fast < l; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums[:slow])
	return slow
}

// 1, 1, 1, 2, 2, 3
// 1, 1, 2, 2, 2, 3

// func main() {
// 	nums := []int{1, 1, 1, 1, 2, 2, 3}
// 	fmt.Println(removeDuplicates(nums))
// }
