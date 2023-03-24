package main

import (
	"fmt"
)

// 243. 最短单词距离
// https://leetcode.cn/problems/shortest-word-distance/
// 方法二
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	index1, index2 := -1, -1
	l := len(wordsDict)
	ans := l
	for i := 0; i < l; i++ {
		if word1 == wordsDict[i] {
			index1 = i
		}
		if word2 == wordsDict[i] {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 && ans > abs(index2-index1) {
			ans = abs(index2 - index1)
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 方法一
//func shortestDistance(wordsDict []string, word1 string, word2 string) int {
//	var word1Position []int
//	var word2Position []int
//	l := len(wordsDict)
//	for i := 0; i < l; i++ {
//		if word1 == wordsDict[i] {
//			word1Position = append(word1Position, i)
//		}
//		if word2 == wordsDict[i] {
//			word2Position = append(word2Position, i)
//		}
//	}
//	shortest := l
//	for i := 0; i < len(word1Position); i++ {
//		for j := 0; j < len(word2Position); j++ {
//			if abs(word1Position[i]-word2Position[j]) < shortest {
//				shortest = abs(word1Position[i] - word2Position[j])
//			}
//		}
//	}
//	return shortest
//}
//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//	return n
//}

func main() {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"
	ret := shortestDistance(wordsDict, word1, word2)
	fmt.Println(ret)
}
