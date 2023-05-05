package main

import (
	"sort"
)

// 1630. 等差子数组
// https://leetcode.cn/problems/arithmetic-subarrays/
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	NumberOfQueries := len(l)
	ans := make([]bool, NumberOfQueries)
	for i := 0; i < NumberOfQueries; i++ {
		if r[i]-l[i] < 2 {
			ans[i] = true
			continue
		}
		arr := make([]int, r[i]-l[i]+1)
		copy(arr, nums[l[i]:r[i]+1])
		ok := query(arr)
		ans[i] = ok
	}
	return ans
}

func query(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// func main() {
// 	nums := []int{4, 6, 5, 9, 3, 7}
// 	l := []int{0, 0, 2}
// 	r := []int{2, 3, 5}
// 	ret := checkArithmeticSubarrays(nums, l, r)
// 	fmt.Println(ret)
// }
