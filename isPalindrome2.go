package main

import (
	"fmt"
)

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/
// 方法二 双指针s
func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		v1, v2 := s[left], s[right]
		for !((v1 >= 48 && v1 <= 57) || (v1 >= 97 && v1 <= 122)) {
			left++
			v1 = s[left]
		}
		for !((v2 >= 48 && v2 <= 57) || (v2 >= 97 && v2 <= 122)) {
			right--
			v2 = s[right]
		}
		fmt.Println(string(v1), string(v2))
		if v1 != v2 {
			return false
		}
		left++
		right--
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
