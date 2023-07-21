package main

// 118. 杨辉三角
// https://leetcode.cn/problems/pascals-triangle/
func generate(numRows int) [][]int {
	var ret [][]int
	ret = append(ret, []int{1})
	for i := 1; i < numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		for j := 1; j < i; j++ {
			tmp = append(tmp, ret[i-1][j-1]+ret[i-1][j])
		}
		tmp = append(tmp, 1)
		ret = append(ret, tmp)
	}
	return ret
}

// func main() {
// 	ret := generate(30)
// 	fmt.Println(ret)
// }
