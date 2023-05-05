package main

// 方法一 递归
func coloredCells(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	return coloredCells(n-1) + int64(4*(n-1))
}

// func main() {
// 	fmt.Println(coloredCells(5))
// }

// 1 1+4*1 5+4*2 13+4*3
// 递归
