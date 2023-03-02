package main

import (
	"fmt"
	"sort"
)

// 2363 合并相似的物品
// https://leetcode.cn/problems/merge-similar-items/description/
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	ret := map[int]int{}
	for _, v := range items1 {
		ret[v[0]] += v[1]
	}
	for _, v := range items2 {
		ret[v[0]] += v[1]
	}
	// 将map装换为二维数组
	var ans [][]int
	for a, b := range ret {
		ans = append(ans, []int{a, b})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

func main() {
	item1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	item2 := [][]int{{3, 1}, {1, 5}}
	ret := mergeSimilarItems(item1, item2)
	fmt.Println(ret)
}
