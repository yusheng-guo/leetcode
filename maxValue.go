package main

// 方法二 优化 迭代(动态规划)
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		for j, x := range g {
			if i > 0 {
				f[i][j] = max(f[i][j], f[i-1][j])
			}
			if j > 0 {
				f[i][j] = max(f[i][j], f[i][j-1])
			}
			f[i][j] += x
		}
	}
	return f[m-1][n-1]
}

// 方法一 递归(动态规划)
// 问题：超时
//func getColumns(input [][]int) [][]int {
//	tmp := make([][]int, len(input))
//	for i := 0; i < len(input); i++ {
//		for j := 1; j < len(input[0]); j++ {
//			tmp[i] = append(tmp[i], input[i][j])
//		}
//	}
//	return tmp
//}
//
//func maxValue(grid [][]int) int {
//	if len(grid) == 1 && len(grid[0]) == 1 {
//		return grid[0][0]
//	}
//	right, down := 0, 0
//	if len(grid) > 1 {
//		down = grid[0][0] + maxValue(grid[1:]) // 向下
//	}
//	if len(grid[0]) > 1 {
//		right = grid[0][0] + maxValue(getColumns(grid)) // 向右
//	}
//	fmt.Println(right, down)
//	if right > down {
//		return right
//	}
//	return down
//}

// func main() {
// 	input := [][]int{
// 		{1, 3, 1},
// 		{1, 5, 1},
// 		{4, 2, 1},
// 	}
// 	ret := maxValue(input)
// 	fmt.Println(ret)
// }
