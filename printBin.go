package main

// 面试题 05.02. 二进制数转字符串
// https://leetcode.cn/problems/bianry-number-to-string-lcci/

func printBin(num float64) string {
	ret := []rune{'0', '.'}
	count := 2 // 计数器
	for count < 32 && num != 0 {
		ret = append(ret, rune(num*2)+'0')
		num = num*2 - float64(int(num*2))
		count++
	}
	if count < 32 {
		return string(ret)
	}
	return "ERROR"
}

// func main() {
// 	ret := printBin(0.625)
// 	fmt.Println(ret)
// }
