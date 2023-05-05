package main

// 119. 杨辉三角 II
// https://leetcode.cn/problems/pascals-triangle-ii/description/
// 方法一
//func getRow(rowIndex int) []int {
//	arr := make([]int, rowIndex+1)
//	arr[0] = 1
//	for i := 1; i < rowIndex+1; i++ {
//		for j := i; j > 0; j-- {
//			arr[j] = arr[j-1] + arr[j]
//		}
//	}
//	return arr
//}

// func main() {
// 	fmt.Println(getRow(5))
// }
