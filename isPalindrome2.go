package main

import (
	"fmt"
	"strings"
)

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/
// 方法二 双指针s
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9')) {
			left++
		}
		for left < right && !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9')) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// 方法一 筛选+判断
//func isPalindrome2(s string) bool {
//	// 大写字符转换为小写字符
//	s = strings.ToLower(s)
//	// 移除所有非字母数字字符
//	var s1 []rune
//	for _, v := range s {
//		if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
//			s1 = append(s1, v)
//		}
//	}
//	// 判断是否是回文串
//	l := len(s1)
//	for i := 0; i < l/2; i++ {
//		if s1[i] != s1[l-i-1] {
//			return false
//		}
//	}
//	return true
//}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome2(s))
}
