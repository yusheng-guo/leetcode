package main

import (
	"fmt"
)

// 1616. 分割两个字符串得到回文串
// https://leetcode.cn/problems/split-two-strings-to-make-palindrome/
// 1 <= a.length, b.length <= 105
// a.length == b.length
// a 和 b 都只包含小写英文字母
// 方法二
func checkPalindromeFormation(a string, b string) bool {
	return check(a, b) || check(b, a)
}

// check 检查 a[:i]+b[i:] 是否为回文串
func check(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	if left >= right {
		return true
	}
	return isIambicString(a[left:right+1]) || isIambicString(b[left:right+1])
}

// isIambicString 检查 s 是否为回文串
func isIambicString(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 方法一：遍历+判断回文数 超时
//func checkPalindromeFormation(a string, b string) bool {
//	l := len(a)
//	if l < 2 {
//		return true
//	}
//	for i := 0; i <= l; i++ {
//		if isPalindrome3(a[:i]+b[i:]) || isPalindrome3(b[:i]+a[i:]) {
//			return true
//		}
//	}
//	return false
//}
//
//func isPalindrome3(s string) bool {
//	left, right := 0, len(s)-1
//	for left < right {
//		if s[left] != s[right] {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

func main() {
	a := "aejbaalflrmkswrydwdkdwdyrwskmrlfqizjezd"
	b := "uvebspqckawkhbrtlqwblfwzfptanhiglaabjea"
	ret := checkPalindromeFormation(a, b)
	fmt.Println(ret)
}
