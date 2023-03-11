package main

import (
	"fmt"
	"strings"
)

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string/
// 方法三 API
func reverseWords(s string) string {
	s = strings.TrimSpace(s) // 去首位空格
	var tmp []rune           // 将多个空格变成一个
	for i, v := range s {
		if v != ' ' {
			tmp = append(tmp, v)
		} else if v == ' ' && s[i-1] != ' ' {
			tmp = append(tmp, ' ')
		}
	}
	arr := strings.Split(string(tmp), " ") // 字符串->单词
	left, right := 0, len(arr)-1           // 反转单词
	for left < right {
		arr[right], arr[left] = arr[left], arr[right]
		left++
		right--
	}
	return strings.Join(arr, " ") // 拼接
}

// 方法二 优化方法一
//func reverseWords(s string) string {
//	s = strings.TrimSpace(s) // 去除前后空字符
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
	s := "  hello   world  "
	ret := reverseWords(s)
	fmt.Println(ret, len(ret))
}
