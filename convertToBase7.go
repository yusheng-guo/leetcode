package main

import (
	"strconv"
)

// 方法三 (方法一优化)
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	tmp := num
	if num < 0 {
		tmp = -num
	}
	ret := ""
	for tmp > 0 {
		ret = strconv.Itoa(tmp%7) + ret
		tmp /= 7
	}
	if num < 0 {
		ret = "-" + ret
	}
	return ret
}

// 方法二
//func convertToBase7(num int) string {
//	return strconv.FormatInt(int64(num), 7)
//}

// 504. 七进制数
// https://leetcode.cn/problems/base-7/
// 方法一
// toBase7 将十进制整数转换为七进制正数
//func toBase7(num int) string {
//	ret := ""
//	for num > 0 {
//		ret = strconv.Itoa(num%7) + ret
//		num /= 7
//	}
//	return ret
//}
//
//func convertToBase7(num int) string {
//	if num == 0 {
//		return "0"
//	} else if num > 0 {
//		return toBase7(num)
//	} else {
//		return "-" + toBase7(-num)
//	}
//}

// func main() {
// 	ret := convertToBase7(100)
// 	fmt.Println(ret)
// }
