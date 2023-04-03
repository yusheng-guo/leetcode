package main

import "fmt"

// LCP 66. 最小展台数量
// https://leetcode.cn/problems/600YaG/description/
// 方法二 方法一优化
func minNumBooths(demand []string) int {
	m0 := map[uint8]int{} // 全局map
	ans := 0
	for i := 0; i < len(demand); i++ {
		m := map[uint8]int{}
		for _, v := range demand[i] {
			m[uint8(v)]++
		}
		for k := range m {
			if m0[k] < m[k] {
				ans += m[k] - m0[k]
				m0[k] = m[k]
			}
		}
	}
	return ans
}

// 方法一
//func minNumBooths(demand []string) int {
//	m0 := map[uint8]int{}
//	ans := 0
//	for _, v := range demand[0] {
//		m0[uint8(v)]++
//		ans++
//	}
//	for i := 1; i < len(demand); i++ {
//		m := map[uint8]int{}
//		for _, v := range demand[i] {
//			m[uint8(v)]++
//		}
//		for k := range m {
//			if m0[k] < m[k] {
//				ans += m[k] - m0[k]
//				m0[k] = m[k]
//			}
//		}
//	}
//	return ans
//}

func main() {
	demand := []string{"acd", "bed", "accd"}
	ret := minNumBooths(demand)
	fmt.Println(ret)
}
