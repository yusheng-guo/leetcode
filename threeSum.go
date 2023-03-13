package main

import "fmt"

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
// 方法一 三重遍历
// isEqual 判断两个三元组是否相等 (不考虑顺序)
func isEqual(arr1, arr2 []int) bool {
	if arr1[0] == arr2[0] && ((arr1[1] == arr2[1] && arr1[2] == arr2[2]) || (arr1[1] == arr2[2] && arr1[2] == arr2[1])) {
		return true
	}
	if arr1[0] == arr2[1] && ((arr1[1] == arr2[2] && arr1[2] == arr2[0]) || (arr1[1] == arr2[0] && arr1[2] == arr2[2])) {
		return true
	}
	if arr1[0] == arr2[2] && ((arr1[1] == arr2[0] && arr1[2] == arr2[1]) || (arr1[1] == arr2[1] && arr1[2] == arr2[0])) {
		return true
	}
	return false
}

func threeSum(nums []int) [][]int {
	var ans [][]int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					equal := false
					for _, tmp := range ans {
						if isEqual(tmp, []int{nums[i], nums[j], nums[k]}) {
							equal = true
							break
						}
					}
					if !equal {
						ans = append(ans, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum(nums)
	fmt.Println(ret)
}
