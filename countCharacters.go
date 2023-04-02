package main

import (
	"fmt"
)

// 1160. 拼写单词
// https://leetcode.cn/problems/find-words-that-can-be-formed-by-characters/
// 方法四 双哈希表 (参考官方题解)
func countCharacters(words []string, chars string) int {
	ans := 0
	lenWords := len(words)
	mChars := map[uint8]int{}
	for _, v := range chars {
		mChars[uint8(v)]++
	}
	for i := 0; i < lenWords; i++ {
		mWord := map[uint8]int{}
		for _, v := range words[i] {
			mWord[uint8(v)]++
		}
		flag := true
		for k := range mWord {
			if mWord[k] > mChars[k] {
				flag = false
				break
			}
		}
		if flag {
			ans += len(words[i])
		}
	}
	return ans
}

// 方法三 遍历 + map
//func countCharacters(words []string, chars string) int {
//	ans := 0
//	lenWords := len(words)
//	for i := 0; i < lenWords; i++ {
//		m := map[uint8]int{}
//		for _, v := range chars {
//			m[uint8(v)]++
//		}
//		for j := 0; j < len(words[i]); j++ {
//			if v := m[words[i][j]]; v != 0 {
//				m[words[i][j]]--
//			} else {
//				break
//			}
//			if j == len(words[i])-1 {
//				ans += len(words[i])
//			}
//		}
//	}
//	return ans
//}

// 方法二 方法一优化
//func countCharacters(words []string, chars string) int {
//	ans := 0
//	lenWords := len(words)
//	for i := 0; i < lenWords; i++ { // 单词
//		tmp := chars
//		lenWord := len(words[i]) // 当前单词的长度
//		isWord := false
//		for j := 0; j < lenWord; j++ { // 字母
//			lenTmp := len(tmp)
//			isAlphabet := false // 字母是否在
//			for k := 0; k < lenTmp; k++ {
//				if words[i][j] == tmp[k] { // words[i][j]在chars中 判断下一个字母
//					isAlphabet = true // 字母存在
//					if k == lenTmp-1 {
//						tmp = tmp[:k]
//					} else {
//						tmp = tmp[:k] + tmp[k+1:]
//					}
//					break
//				}
//			}
//			if !isAlphabet { // words[i][j]不在chars中 判断下一个单词
//				break
//			}
//			if j == lenWord-1 {
//				isWord = true
//			}
//		}
//		if isWord {
//			ans += lenWord
//		}
//	}
//	return ans
//}

// 方法一 暴力枚举 三层循环
//func countCharacters(words []string, chars string) int {
//	ans := 0
//	for i := 0; i < len(words); i++ { // 单词
//		tmp := chars
//		isWord := false
//		for j := 0; j < len(words[i]); j++ { // 字母
//			isAlphabet := false // 字母是否在
//			for k := 0; k < len(tmp); k++ {
//				if words[i][j] == tmp[k] { // words[i][j]在chars中 判断下一个字母
//					isAlphabet = true // 字母存在
//					if k == len(tmp)-1 {
//						tmp = tmp[:k]
//					} else {
//						tmp = tmp[:k] + tmp[k+1:]
//					}
//					break
//				}
//			}
//			if !isAlphabet { // words[i][j]不在chars中 判断下一个单词
//				break
//			}
//			if j == len(words[i])-1 {
//				isWord = true
//			}
//		}
//		if isWord {
//			ans += len(words[i])
//		}
//	}
//	return ans
//}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	ret := countCharacters(words, chars)
	fmt.Println(ret)
}
