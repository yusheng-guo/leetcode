package main

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
// 方法二 双指针
// fast 快指针 找出不同的元素
// slow
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 { // 0个或1个元素
		return l
	}
	slow := 1 //
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 方法一 递归 不考虑空间问题
//func removeDuplicates(nums []int) int {
//	l := len(nums)
//	fmt.Println(nums)
//	if l < 2 {
//		return l
//	}
//	if nums[0] == nums[1] {
//		return removeDuplicates(nums[1:l])
//	}
//	return removeDuplicates(nums[1:l]) + 1
//}

/*
0, 0, 1, 1, 1, 2, 2, 3, 3, 4
0, 1, 1, 1, 2, 2, 3, 3, 4
*/
// func main() {
// 	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	//nums := []int{1, 1, 2}
// 	fmt.Println(removeDuplicates(nums))
// }
