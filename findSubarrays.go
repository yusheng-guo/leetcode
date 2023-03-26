package main

import "fmt"

// 2395. 和相等的子数组
// https://leetcode.cn/problems/find-subarrays-with-equal-sum/
func findSubarrays(nums []int) bool {
	l := len(nums)
	adds := map[int]struct{}{}
	for i := 0; i < l-1; i++ {
		if _, ok := adds[nums[i]+nums[i+1]]; ok {
			return true
		}
		adds[nums[i]+nums[i+1]] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ret := findSubarrays(nums)
	fmt.Println(ret)
}
