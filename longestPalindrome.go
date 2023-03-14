package main

import (
	"fmt"
)

// 409. 最长回文串
// https://leetcode.cn/problems/longest-palindrome/

// 方法一 双重循环 暴力解法
// 问题 超出时间限制
// is 判断是否为回文串
func is(s string) bool {
	reverse := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return s == string(reverse)
}

func longestPalindrome(s string) string {
	l := len(s)
	max := 0
	left, right := 0, 0
	for i := 0; i < l; i++ {
		for j := l; j-i > max; j-- {
			if is(s[i:j]) {
				max = j - i
				left, right = i, j
			}
		}
	}
	return s[left:right]
}

func main() {
	s := "babad"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}
