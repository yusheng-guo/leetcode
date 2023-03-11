package main

import (
	"fmt"
)

// 415. 字符串相加
// https://leetcode.cn/problems/add-strings/
// 方法二
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	arr := make([]byte, maxInt(l1, l2)+2)

	//if l1 == l2 {
	//	if num1[0]+num2[0]-'0'*2 > 9 {
	//		arr = make([]byte, l1+1)
	//	} else {
	//		arr = make([]byte, l1)
	//	}
	//} else if l1 > l2 {
	//	arr = make([]byte, l1)
	//} else {
	//	arr = make([]byte, l2)
	//}
	p1, p2, p3 := l1-1, l2-1, len(arr)-1
	carry := false // 进位
	for p1 >= 0 && p2 >= 0 {
		var tmp uint8       // 位和
		if carry == false { // 没有进位
			tmp = num1[p1] + num2[p2] - '0'
		} else { // 有进位
			tmp = num1[p1] + num2[p2] - '0' + 1
		}
		if tmp > '9' { // 产生进位
			arr[p3] = tmp - 10
			carry = true
		} else { // 没有进位
			arr[p3] = tmp
		}
		p1--
		p2--
		p3--
	}
	if p1 < 0 && p2 < 0 {
		if carry == true {
			arr[p3] = '1'
		}
		return string(arr)
	}
	if p1 < 0 { // num2剩余
		for p2 >= 0 {
			if carry == true {
				arr[p3] = num2[p2] + 1
			} else {
				arr[p3] = num2[p2]
			}
			p2--
			p3--
		}
		return string(arr)
	}
	if p2 < 0 { // num1剩余
		for p1 >= 0 {
			if carry == true {
				arr[p3] = num1[p2] + 1
			} else {
				arr[p3] = num1[p2]
			}
			p1--
			p3--
		}
		return string(arr)
	}
	return string(arr)
}

// 方法一 API
// 问题：不能计算大数
//func addStrings(num1 string, num2 string) string {
//	a, _ := strconv.Atoi(num1)
//	b, _ := strconv.Atoi(num2)
//	return strconv.Itoa(a + b)
//}

func main() {
	ret := addStrings("3876620623801494171", "6529364523802684779")
	fmt.Println(ret)
}
