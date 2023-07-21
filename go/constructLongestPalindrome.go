package main

func constructLongestPalindrome(s string) int {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	ans := 0
	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else {
			ans += v - 1
		}
	}
	if ans == len(s) {
		return ans
	}
	return ans + 1
}

// func main() {
// 	s := "abccccdd"
// 	ret := constructLongestPalindrome(s)
// 	fmt.Println(ret)
// }
