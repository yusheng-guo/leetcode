package main

import (
	"sort"
)

// 2418. 按身高排序
// https://leetcode.cn/problems/sort-the-people/
// 方法二 索引数组
func sortPeople(names []string, heights []int) []string {
	n := len(heights)
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return heights[idxs[i]] > heights[idxs[j]]
	})
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = names[idxs[i]]
	}
	return ret
}

// 方法一 冒泡排序
//func sortPeople(names []string, heights []int) []string {
//	for i := 0; i < len(heights); i++ {
//		for j := i; j < len(heights); j++ {
//			if heights[i] < heights[j] {
//				heights[i], heights[j] = heights[j], heights[i]
//				names[i], names[j] = names[j], names[i]
//			}
//		}
//	}
//	return names
//}

// func main() {
// 	names := []string{"a", "b", "c", "d", "e"}
// 	heights := []int{1, 5, 4, 2, 3}
// 	ret := sortPeople(names, heights)
// 	fmt.Println(ret)
// }

//[b c e d a]
//[e a b d c]
