package main

import (
	"fmt"
)

// 2379. 得到 K 个黑块的最少涂色次数
// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// 方法三 滑动窗口优化 注：复杂度不变
func minimumRecolors(blocks string, k int) int {
	left, right, l := 0, 0, len(blocks)
	white := 0
	min := l
	for ; right < l; right++ {
		if blocks[right] == 'W' {
			white++
		}
		if right < k-1 {
			continue
		}
		if white < min {
			min = white
		}
		if blocks[left] == 'W' {
			white--
		}
		left++
	}
	return min
}

// 方法二 滑动窗口
//func minimumRecolors(blocks string, k int) int {
//	left, right, l := 0, 0, len(blocks)
//	white := 0 // 白色块个数
//	for ; right < k; right++ {
//		if blocks[right] == 'W' {
//			white++
//		}
//	}
//	min := white
//	for ; right < l; right++ {
//		if blocks[left] == 'W' {
//			white--
//		}
//		if blocks[right] == 'W' {
//			white++
//		}
//		if white < min {
//			min = white
//		}
//		left++
//	}
//	return min
//}

// 方法一 双循环 时间复杂度 O(len(blocks)*k)
//func minimumRecolors(blocks string, k int) int {
//	min := math.MaxInt
//	l := len(blocks) // blocks字符串长度
//	for i := 0; i < l-k+1; i++ {
//		tmp := 0
//		for j := 0; j < k; j++ {
//			if blocks[i+j] == 'W' {
//				tmp++
//			}
//		}
//		if tmp < min {
//			min = tmp
//		}
//	}
//	return min
//}

func main() {
	blocks := "BWWWBB"
	k := 6
	ret := minimumRecolors(blocks, k)
	fmt.Println(ret)
}
