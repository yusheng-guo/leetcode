package main

// 1653. 使字符串平衡的最少删除次数
// https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/

// 方法二 优化 单循环
func minimumDeletions(s string) int {
	righta := 0 // 分割线右边a的个数
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			righta++
		}
	}
	min := righta // 最小分割次数
	leftb := 0    //分割线左边b的个数
	// 分割线位置
	for j := 0; j < len(s); j++ {
		if s[j] == 'a' {
			righta--
		} else {
			leftb++
		}
		if min > righta+leftb {
			min = righta + leftb
		}
	}
	return min
}

// 方法一 双循环
// 问题：超时
//func minimumDeletions(s string) int {
//	min := math.MaxInt // 最小删除次数
//	for i := 0; i <= len(s); i++ {
//		// 分割线为 i
//		leftb := 0  // i之前b的数量
//		righta := 0 // i之后a的数量
//		for j := 0; j < i; j++ {
//			if s[j] == 'b' {
//				leftb++
//			}
//		}
//		for j := i; j < len(s); j++ {
//			if s[j] == 'a' {
//				righta++
//			}
//		}
//		if leftb+righta < min {
//			min = leftb + righta
//		}
//	}
//	return min
//}

// func main() {
// 	s := "aaabbbabbabaabbbb"
// 	ret := minimumDeletions(s)
// 	fmt.Println(ret)
// }
