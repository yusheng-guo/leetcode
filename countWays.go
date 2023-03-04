package main

import "fmt"

// ???
func countWays(ranges [][]int) int {
	l := len(ranges)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if ranges[i][0] <= ranges[j][1] && ranges[i][1] >= ranges[j][0] {
				// 两个区间有交集
			}
		}
	}
	return 0
}

func main() {
	ranges := [][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}}
	ret := countWays(ranges)
	fmt.Println(ret)
}
