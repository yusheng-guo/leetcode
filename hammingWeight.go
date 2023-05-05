package main

// 191. 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/description/
// Brian Kernighan 算法
func hammingWeight(num uint32) (count int) {
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}

// func main() {
// 	ret := hammingWeight(11)
// 	fmt.Println(ret)
// }
