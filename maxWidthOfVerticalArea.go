package main

import (
	"fmt"
	"sort"
)

// 1637. 两点之间不包含任何点的最宽垂直区域
// https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points/
// 方法二 优化空间复杂度
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ret := 0
	for i := 1; i < len(points); i++ {
		if ret < points[i][0]-points[i-1][0] {
			ret = points[i][0] - points[i-1][0]
		}
	}
	return ret
}

// 方法一
//func maxWidthOfVerticalArea(points [][]int) int {
//	l := len(points)
//	arr := make([]int, l)
//	for i := 0; i < l; i++ {
//		arr[i] = points[i][0]
//	}
//	sort.Ints(arr)
//	ans := 0
//	for i := 1; i < len(arr); i++ {
//		if ans < arr[i]-arr[i-1] {
//			ans = arr[i] - arr[i-1]
//		}
//	}
//	return ans
//}

func main() {
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	ret := maxWidthOfVerticalArea(points)
	fmt.Println(ret)
}
