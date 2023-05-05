package main

import (
	"bytes"
)

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 方法三
//func lengthOfLongestSubstring(s string) int {
//	// 哈希集合，记录每个字符是否出现过
//	m := map[byte]int{}
//	n := len(s)
//	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
//	rk, ans := -1, 0
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			// 左指针向右移动一格，移除一个字符
//			delete(m, s[i-1])
//		}
//		for rk+1 < n && m[s[rk+1]] == 0 {
//			// 不断地移动右指针
//			m[s[rk+1]]++
//			rk++
//		}
//		// 第 i 到 rk 个字符是一个极长的无重复字符子串
//		if ans < rk-i+1 {
//			ans = rk - i + 1
//		}
//	}
//	return ans
//}

// 方法二 哈希表 滑动窗口
//func lengthOfLongestSubstring(s string) int {
//	l := len(s)
//	right, ans := 0, 0
//	m := map[byte]struct{}{}
//	for i := 0; i < l; i++ {
//		if i != 0 {
//			delete(m, s[i-1])
//		}
//		for right < l {
//			if _, ok := m[s[right]]; !ok {
//				m[s[right]] = struct{}{}
//				right++
//			} else {
//				break
//			}
//		}
//		if ans < len(m) {
//			ans = len(m)
//		}
//	}
//	return ans
//}

// 方法一 思想 滑动窗口
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	ans := 0
	if l < 2 {
		return l
	}
	left, right := 0, 1
	for left < right && right < l {
		idx := bytes.IndexByte([]byte((s[left:right])), s[right])
		if idx != -1 {
			left += idx + 1
		}
		if ans < right-left+1 {
			ans = right - left + 1
		}
		right++
	}
	return ans
}

// func main() {
// 	s := "abcabc"
// 	ret := lengthOfLongestSubstring(s)
// 	fmt.Println(ret)
// }
