package main

import (
	"bytes"
	"fmt"
)

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 思想 滑动窗口
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

func main() {
	s := "abcabcbb"
	ret := lengthOfLongestSubstring(s)
	fmt.Println(ret)
}
