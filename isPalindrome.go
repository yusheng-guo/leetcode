package main

import "fmt"

// 回文数
// https://leetcode.cn/problems/palindrome-number/
//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	str := []byte(strconv.Itoa(x))
//	//for i := 0; i < len(str)-1; i++ {
//	//	for j := i + 1; j < len(str); j++ {
//	//		str[i], str[j] = str[j], str[i]
//	//	}
//	//}
//	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
//		str[i], str[j] = str[j], str[i]
//	}
//	i, _ := strconv.Atoi(string(str))
//	return x == i
//}

//func isPalindrome(x int) bool {
//	tmp := strconv.Itoa(x)
//	str := []byte(tmp)
//	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
//		str[i], str[j] = str[j], str[i]
//	}
//	return tmp == string(str)
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num, ret := x, 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	return num == ret
}

func main() {
	ret := isPalindrome(12321)
	fmt.Println(ret)
}
