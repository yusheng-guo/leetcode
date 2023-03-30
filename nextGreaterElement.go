package main

import "fmt"

// 496. 下一个更大元素 I
// https://leetcode.cn/problems/maxWidthOfVerticalAreanext-greater-element-i/
// 方法二 单调栈+哈希表
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	ans := make([]int, l1)
	return ans
}

// 方法一 暴力枚举
//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	l1, l2 := len(nums1), len(nums2)
//	ans := make([]int, l1)
//	for i := 0; i < l1; i++ {
//		flag1, flag2 := false, false
//		for j := 0; j < l2; j++ {
//			if flag1 == true && nums1[i] < nums2[j] {
//				ans[i] = nums2[j]
//				flag2 = true
//				break
//			}
//			if nums1[i] == nums2[j] {
//				fmt.Println(i, j)
//				flag1 = true
//			}
//		}
//		if !flag2 {
//			ans[i] = -1
//		}
//	}
//	return ans
//}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	ret := nextGreaterElement(nums1, nums2)
	fmt.Println(ret)
}
