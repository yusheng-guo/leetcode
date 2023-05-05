package main

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome1(s string) string {
	l := len(s)
	ret := string(s[0])
	for i := 0; i < l; i++ {
		// "***xx***"形式
		if i < l-1 && s[i] == s[i+1] {
			j := 1
			for i-j >= 0 && i+j+1 < l && s[i-j] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2 { // 更新ret
				ret = s[i-j+1 : i+j+1]
			}
		}

		// "***xyx***" 形式
		if i > 0 && i < l-1 && s[i-1] == s[i+1] {
			j := 1
			for i-j-1 >= 0 && i+j+1 < l && s[i-j-1] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2+1 { // 更新ret
				ret = s[i-j : i+j+1]
			}
		}
	}
	return ret
}

// func main() {
// 	s := "bababad"
// 	ret := longestPalindrome1(s)
// 	fmt.Println(ret)
// }
