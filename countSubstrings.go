package main

import "fmt"

// 1638. 统计只差一个字符的子串数目
// https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/
// 方法一 遍历 三次循环
func countSubstrings(s string, t string) int {
	lens, lent := len(s), len(t)
	ret := 0
	for i := 0; i < lens; i++ {
		for j := 0; j < lent; j++ {
			k := 0
			diff := 0
			for k < min(lens-i, lent-j) && diff < 2 {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff == 1 {
					ret++
				}
				k++
			}
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "aba"
	t := "baba"
	ret := countSubstrings(s, t)
	fmt.Println(ret)
}
