package main

import (
	"fmt"
	"strings"
)

// 方法二 优化方法一
func reverseWords(s string) string {
	s = strings.TrimSpace(s) // 去除前后空字符
	fmt.Println(s)
	if len(s) < 3 {
		return s
	}
	ret := ""
	i := len(s) - 1
	for i > 0 {
		// 找到单词的最后一个字符
		for i >= 0 && s[i] == ' ' {
			i--
		}
		last := i + 1
		// 找到单词的第一个字符
		for i >= 0 && s[i] != ' ' {
			i--
		}
		first := i + 1
		ret = ret + " " + s[first:last]
	}
	return ret[1:]
}

// 方法一
//func revers//	// 去除前后空字符
////	left, right := 0, len(s)-1
////	for ; right >= 0 && s[right] == ' '; right-- {
////	}
////	for ; left < right && s[left] == ' '; left++ {
////	}
////	s = s[left : right+1]eWords(s string) string {

//	fmt.Println(s)
//	if len(s) < 3 {
//		return s
//	}
//	ret := ""
//	i := len(s) - 1
//	for i > 0 {
//		// 找到单词的最后一个字符
//		for i >= 0 && s[i] == ' ' {
//			i--
//		}
//		last := i + 1
//		// 找到单词的第一个字符
//		for i >= 0 && s[i] != ' ' {
//			i--
//		}
//		first := i + 1
//		ret = ret + " " + s[first:last]
//	}
//	return ret[1:]
//}

func main() {
	//s := "the sky is blue"
	//s := "hello world!"
	s := "a"
	ret := reverseWords(s)
	fmt.Println(ret, len(ret))
}
