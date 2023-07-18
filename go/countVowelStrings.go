package main

// 1641. 统计字典序元音字符串的数目
// https://leetcode.cn/problems/count-sorted-vowel-strings/
// 方法三 动态规划 优化方法二 一数组
func countVowelStrings(n int) int {
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		dp[0] = 1
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}
	sum := 0
	for _, v := range dp {
		sum += v
	}
	return sum
}

// 方法二 动态规划 二维数组
//func countVowelStrings(n int) int {
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, 5)
//	}
//	dp[0] = []int{1, 1, 1, 1, 1}
//	sum := 5
//	for i := 1; i < n; i++ {
//		dp[i][0] = sum
//		for j := 1; j < 5; j++ {
//			dp[i][j] = dp[i][j-1] - dp[i-1][j-1]
//			sum += dp[i][j]
//		}
//	}
//	return sum
//}

// 方法一 排列组合 C((n+4), 4)
//func countVowelStrings(n int) int {
//	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
//}

// func main() {
// 	ret := countVowelStrings(2)
// 	fmt.Println(ret)
// }
