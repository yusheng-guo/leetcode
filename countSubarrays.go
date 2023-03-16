package main

import "fmt"

// 2488. 统计中位数为 K 的子数组
// https://leetcode.cn/problems/count-subarrays-with-median-k/
func countSubarrays(nums []int, k int) int {
	kIndex := -1
	for i, num := range nums {
		if num == k {
			kIndex = i
			break
		}
	}
	fmt.Println(kIndex)
	ans := 0
	counts := map[int]int{}
	counts[0] = 1
	sum := 0
	for i, num := range nums {
		sum += sign(num - k)
		if i < kIndex {
			counts[sum]++
		} else {
			prev0 := counts[sum]
			prev1 := counts[sum-1]
			ans += prev0 + prev1
		}
	}
	return ans
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}

func main() {
	nums := []int{3, 2, 1, 4, 5}
	k := 4
	ret := countSubarrays(nums, k)
	fmt.Println(ret)
}
