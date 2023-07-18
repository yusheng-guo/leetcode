package main

import "fmt"

// 矩阵中的局部最大值
// https://leetcode.cn/problems/largest-local-values-in-a-matrix/
func largestLocal(grid [][]int) [][]int {
	r := len(grid)    // 行数
	c := len(grid[0]) // 列数
	ret := make([][]int, r-2)
	for i := 0; i < r-2; i++ {
		for j := 0; j < c-2; j++ {
			max := 0
			fmt.Println(i, j)
			for k := 0; k < 3; k++ {
				for z := 0; z < 3; z++ {
					if grid[k+i][z+j] > max {
						max = grid[k+i][z+j]
					}
				}
			}
			fmt.Println(max)
			ret[i] = append(ret[i], max)
		}
	}
	return ret
}

// func main() {
// 	grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
// 	ret := largestLocal(grid)
// 	fmt.Println(ret)
// }
