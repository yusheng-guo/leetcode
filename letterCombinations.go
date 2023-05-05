package main

import "fmt"

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	var ret []string
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	count := 1
	// 计算个数
	for _, v := range digits {
		count *= len(m[byte(v)])
	}
	fmt.Println(count)

	return ret
}

// func main() {
// 	digits := "23"
// 	ret := letterCombinations(digits)
// 	fmt.Println(ret)
// }
