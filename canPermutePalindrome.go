package main

import "fmt"

// 266. 回文排列
// https://leetcode.cn/problems/palindrome-permutation/description/
// 方法二
func canPermutePalindrome(s string) bool {
	m := map[byte]struct{}{}
	for _, v := range s {
		_, ok := m[byte(v)]
		if ok {
			delete(m, byte(v))
		} else {
			m[byte(v)] = struct{}{}
		}
	}
	return len(m) < 2
}

// 方法一
//func canPermutePalindrome(s string) bool {
//	m := map[byte]int{}
//	for _, v := range s {
//		m[byte(v)]++
//	}
//	count := 0
//	for _, v := range m {
//		if v%2 != 0 {
//			count++
//		}
//	}
//	return count < 2
//}

func main() {
	ret := canPermutePalindrome("abdg")
	fmt.Println(ret)
}
