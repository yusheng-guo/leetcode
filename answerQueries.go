package main

import (
	"fmt"
	"sort"
)

// 2389. 和有限的最长子序列
// https://leetcode.cn/problems/longest-subsequence-with-limited-sum/
func answerQueries(nums []int, queries []int) []int {
	l := len(nums)
	sort.Ints(nums)
	var arr []int
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		arr = append(arr, sum)
	}
	var ans []int
	for i := 0; i < len(queries); i++ {
		z := l
		if arr[0] > queries[i] {
			z = 0
		}
		for j := 1; j < l; j++ {
			if arr[j] > queries[i] && arr[j-1] <= queries[i] {
				z = j
				break
			}
		}
		ans = append(ans, z)
	}
	return ans
}

func main() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	ret := answerQueries(nums, queries)
	fmt.Println(ret)
}
