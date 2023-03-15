package main

import "fmt"

// 168. Excel表列名称
// https://leetcode.cn/problems/excel-sheet-column-title/
// 二十六进制
func convertToTitle(columnNumber int) string {
	ret := ""
	for columnNumber > 0 {
		columnNumber--
		ret = string(rune((columnNumber%26)+'A')) + ret
		columnNumber /= 26
	}
	return ret
}

func main() {
	columnNumber := 27
	ret := convertToTitle(columnNumber)
	fmt.Println(ret)
}
