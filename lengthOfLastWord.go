package main

import "fmt"

// 58. 最后一个单词的长度
// https://leetcode.cn/problems/length-of-last-word/
// 注意：临界判断 比如：""  " " "a"
// 思想：反向遍历 分段遍历
// 方法二
func lengthOfLastWord(s string) int {
	// 找到最后一个单词的最后一个字符
	i := len(s) - 1
	n := 0
	for s[i] == ' ' {
		i--
	}
	// 找到最后一个单词的第一个字符
	for i >= 0 && s[i] != ' ' {
		i--
		n++
	}
	return n
}

// 方法一
//func lengthOfLastWord(s string) int {
//	var p1, p2 int
//	l := len(s)
//	if l == 0 {
//		return 0
//	}
//	for i := l - 1; i >= 0; i-- {
//		if s[i] != ' ' && p1 == 0 {
//			p1 = i
//		}
//		if s[i] == ' ' && p1 != 0 && p2 == 0 {
//			p2 = i
//			return p1 - p2
//		}
//	}
//	return p1 - p2 + 1
//}

func main() {
	ret := lengthOfLastWord(" abcs   ")
	fmt.Println(ret)
}
