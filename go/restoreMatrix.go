package main

// 1605. 给定行和列的和求可行矩阵
// https://leetcode.cn/problems/find-valid-matrix-given-row-and-column-sums/

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	// 创建二维数组
	m, n := len(rowSum), len(colSum) // 行 列
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	i, j := 0, 0
	for i < m && j < n {
		v := min(rowSum[i], colSum[j])
		matrix[i][j] = v
		rowSum[i] -= v
		colSum[j] -= v
		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return matrix
}

// func main() {
// 	rowSum := []int{3, 8}
// 	colSum := []int{4, 7}
// 	ret := restoreMatrix(rowSum, colSum)
// 	fmt.Println(ret)
// }
