package main

import "fmt"

// 面试题 17.05. 字母与数字
// https://leetcode.cn/problems/find-longest-subarray-lcci/description/
// leetcode 官方
//func findLongestSubarray(array []string) []string {
//	indices := map[int]int{}
//	sum := 0
//	startIndex := -1
//	maxLength := 0
//	indices[0] = -1
//	for i, s := range array {
//		if s[0] >= '0' && s[0] <= '9' {
//			sum++
//		} else {
//			sum--
//		}
//		if firstIndex, ok := indices[sum]; ok {
//			if i-firstIndex > maxLength {
//				maxLength = i - firstIndex
//				startIndex = firstIndex + 1
//			}
//		} else {
//			indices[sum] = i
//		}
//	}
//	if maxLength == 0 {
//		return []string{}
//	}
//	return array[startIndex : startIndex+maxLength]
//}

func findLongestSubarray(array []string) []string {
	return []string{}
}

func main() {
	array := []string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}
	ret := findLongestSubarray(array)
	fmt.Println(ret)

	test := "9"
	fmt.Println(test)
	fmt.Println(test[0])
	//'A' 65
	//'Z' 90
	//'0' 48
	//'9'
}
