package main

import "fmt"

// 剑指 Offer II 002. 二进制加法
// https://leetcode.cn/problems/JFETK5/description/
func addBinary(a string, b string) string {
	ans := ""
	for l1, l2, carry := len(a)-1, len(b)-1, uint8(0); l1 >= 0 || l2 >= 0 || carry == 1; l1, l2 = l1-1, l2-1 {
		var tmp1, tmp2 uint8 = 0, 0
		if l1 >= 0 {
			tmp1 = a[l1] - '0'
		}
		if l2 >= 0 {
			tmp2 = b[l2] - '0'
		}
		ret := tmp1 + tmp2 + carry
		ans = string(ret%2+'0') + ans
		if ret > 1 {
			carry = 1
		} else {
			carry = 0
		}
	}
	return ans
}

func main() {
	a := "1010"
	b := "1011"
	ret := addBinary(a, b)
	fmt.Println(ret)
}
