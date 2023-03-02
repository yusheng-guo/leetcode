package main

import (
	"fmt"
)

// 找出字符串的可整除数组
func divisibilityArray(word string, m int) []int {
	ans := []int{}
	var num int = 0
	for _, char := range word {
		num = num*10 + int(char-'0')
		if num%m == 0 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
		num %= m
	}
	return ans
}

//func divisibilityArray(word string, m int) []int {
//	l := len(word)
//	arr := []rune(word)
//	ans := make([]int, l)
//	for i := 0; i < l; i++ {
//		tmp, _ := strconv.Atoi(string(arr[0 : i+1]))
//		fmt.Println(tmp)
//		if tmp%m == 0 {
//			ans[i] = 1
//		} else {
//			ans[i] = 0
//		}
//	}
//	return ans
//}

func main() {
	ret := divisibilityArray("998244353", 3)
	fmt.Println(ret)
}
