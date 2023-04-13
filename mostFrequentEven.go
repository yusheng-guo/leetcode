package main

import (
	"fmt"
)

// 2404. 出现最频繁的偶数元素
// https://leetcode.cn/problems/most-frequent-even-element/
// 方法二 优化方法一
func mostFrequentEven(num []int) int {
	frequent := map[int]int{}
	//frequent := make(map[int]int, len(num))
	for _, v := range num {
		if v&1 == 0 {
			frequent[v]++
		}
	}
	res, most := -1, 0
	for k, v := range frequent {
		if most < v || most == v && k < res {
			res = k
			most = v
		}
	}
	return res
}

// 方法一
//func mostFrequentEven(num []int) (minMost int) {
//	frequent := map[int]int{}
//	for _, v := range num {
//		if v%2 == 0 {
//			frequent[v]++
//		}
//	}
//	// 获取最高的频率
//	most := 0
//	for _, v := range frequent {
//		if most < v {
//			most = v
//		}
//	}
//	// 出现最高的频率的元素
//	var arr []int
//	for k, v := range frequent {
//		if v == most {
//			arr = append(arr, k)
//		}
//	}
//	if len(arr) == 0 {
//		return -1
//	}
//	// 其中的最小值
//	minMost = math.MaxInt
//	for _, v := range arr {
//		if v < minMost {
//			minMost = v
//		}
//	}
//	return
//}

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	ret := mostFrequentEven(nums)
	fmt.Println(ret)
}
