package main

import (
	"fmt"
	"math"
)

// 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	l := len(strs) // 字符串数组长度
	if l <= 1 {
		return strs[0]
	}
	minLength := math.MaxInt32 // 最短字符串长度
	s := ""                    // 字符串前缀
	for i := 0; i < l; i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}
	fmt.Println(l, minLength, s)
	for k := 0; k < minLength; k++ {
		s = strs[0][0 : k+1]
		for j := 1; j < l; j++ {
			if s != strs[j][0:k+1] {
				return s[0:k]
			}
		}
	}
	return s
}

// func main() {
// 	//strs := []string{"flower", "flow", "flight"}
// 	strs := []string{"ab", "a"}
// 	ret := longestCommonPrefix(strs)
// 	fmt.Println(ret)
// }
