package main

import (
	"fmt"
)

// 409. 最长回文串
// https://leetcode.cn/problems/longest-palindrome/

// 方法一 双重循环
// is 判断是否为回文串
func is(s string) bool {
	reverse := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return s == string(reverse)
}

func longestPalindrome(s string) int {
	l := len(s)
	max := 0
	for i := 0; i < l; i++ {
		for j := l - 1; j > i; j-- {
			if is(s[i:j+1]) && len(s[i:j+1]) > max {
				max = len(s[i : j+1])
			}
		}
	}
	return max
}

func main() {
	s := "abccccdd"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}
