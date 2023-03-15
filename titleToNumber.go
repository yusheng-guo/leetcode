package main

import "fmt"

// 171. Excel 表列序号
// https://leetcode.cn/problems/excel-sheet-column-number/description/
func titleToNumber(columnTitle string) int {
	ret := 0
	for i := 0; i < len(columnTitle); i++ {
		fmt.Println(columnTitle[i], ret)
		ret = ret*26 + int((columnTitle[i]-'A')+1)
	}
	return ret
}

func main() {
	columnTitle := "FXSHRXW"
	ret := titleToNumber(columnTitle)
	fmt.Println(ret)
}
