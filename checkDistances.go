package main

// 2399. 检查相同字母间的距离
// https://leetcode.cn/problems/check-distances-between-same-letters/
// 方法三 Leetcode官方题解
//func checkDistances(s string, distance []int) bool {
//	firstIndex := make([]int, 26)
//	for i := 0; i < len(s); i++ {
//		idx := s[i] - 'a'
//		if firstIndex[idx] != 0 && i-firstIndex[idx] != distance[idx] {
//			return false
//		}
//		firstIndex[idx] = i + 1
//	}
//	return true
//}

// 方法二 方法一优化
func checkDistances(s string, distance []int) bool {
	for i, v := range s {
		if distance[v-'a'] == -1 { // 判断是否已判断
			continue
		}
		dis := distance[v-'a']
		if i+dis >= len(s)-1 || uint8(v) != s[i+dis+1] {
			return false
		}
		distance[v-'a'] = -1 // 标记为已判断
	}
	return true
}

// 方法一
//func checkDistances(s string, distance []int) bool {
//	for i, v := range s {
//		dis := distance[v-'a']
//		if !((i+dis < len(s)-1 && uint8(v) == s[i+dis+1]) || (i-dis > 0 && uint8(v) == s[i-dis-1])) {
//			return false
//		}
//	}
//	return true
//}

// func main() {
// 	s := "abaccb"
// 	distance := []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	ret := checkDistances(s, distance)
// 	fmt.Println(ret)
// }
