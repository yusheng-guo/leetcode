package main

import (
	"fmt"
)

// 1033. 移动石子直到连续
// https://leetcode.cn/problems/moving-stones-until-consecutive/
func numMovesStones(a int, b int, c int) []int {
	// 排序
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	max := c - a - 2
	min := 2
	if a+1 == b || b+1 == c || a+2 == b || b+2 == c {
		min = 1
	}
	if a+1 == b && b+1 == c {
		min = 0
	}
	return []int{min, max}
}

func main() {
	a := 1
	b := 2
	c := 5
	ret := numMovesStones(a, b, c)
	fmt.Println(ret)
}
