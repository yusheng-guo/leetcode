package main

import (
	"sort"
)

// 1626. 无矛盾的最佳球队
// https://leetcode.cn/problems/best-team-with-no-conflicts/
// 思路 动态规划
// 排序：按分数排序，分数相同则按年龄排序
func bestTeamScore(scores []int, ages []int) int {
	// 合并
	l := len(scores)
	people := make([][2]int, l)
	for i := 0; i < l; i++ {
		people[i][0] = scores[i]
		people[i][1] = ages[i]
	}
	// 排序 先按分数排序 分数相同按年龄排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		return people[i][1] < people[j][1]
	})
	arr := make([]int, l)
	maxScore := 0
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if people[i][1] >= people[j][1] {
				arr[i] = max(arr[i], arr[j])
			}
		}
		arr[i] += people[i][0]
		maxScore = max(maxScore, arr[i])
	}
	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode 官方题解
//func bestTeamScore(scores []int, ages []int) int {
//	n := len(scores)
//	people := make([][]int, n)
//	for i := range scores {
//		people[i] = []int{scores[i], ages[i]}
//	}
//	sort.Slice(people, func(i, j int) bool {
//		if people[i][0] < people[j][0] {
//			return true
//		} else if people[i][0] > people[j][0] {
//			return false
//		}
//		return people[i][1] < people[j][1]
//	})
//	dp := make([]int, n)
//	res := 0
//	for i := 0; i < n; i++ {
//		for j := 0; j < i; j++ {
//			if people[j][1] <= people[i][1] {
//				dp[i] = max(dp[i], dp[j])
//			}
//		}
//		dp[i] += people[i][0]
//		res = max(res, dp[i])
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}

// func main() {
// 	scores := []int{4, 5, 6, 5}
// 	ages := []int{2, 1, 2, 1}
// 	ret := bestTeamScore(scores, ages)
// 	fmt.Println(ret)
// }
