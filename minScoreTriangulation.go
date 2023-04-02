package main

import (
	"fmt"
	"math"
)

// 1039. 多边形三角剖分的最低得分
// https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/

func minScoreTriangulation(values []int) int {
	memo := make(map[int]int)
	n := len(values)
	var dp func(int, int) int
	dp = func(i int, j int) int {
		if i+2 > j {
			return 0
		}
		if i+2 == j {
			return values[i] * values[i+1] * values[j]
		}
		key := i*n + j
		if _, ok := memo[key]; !ok {
			minScore := math.MaxInt32
			for k := i + 1; k < j; k++ {
				minScore = min(minScore, values[i]*values[k]*values[j]+dp(i, k)+dp(k, j))
			}
			memo[key] = minScore
		}
		return memo[key]
	}
	return dp(0, n-1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	values := []int{1, 2, 3}
	ret := minScoreTriangulation(values)
	fmt.Println(ret)
}
