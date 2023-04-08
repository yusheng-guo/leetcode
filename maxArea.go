package main

import (
	"fmt"
)

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/
// 思想 双指针
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0
	for left < right {
		// 获取高度
		h := height[left]
		if height[left] > height[right] {
			h = height[right]
		}
		// 计算容积
		area := (right - left) * h
		if ret < area {
			ret = area
		}
		// 移动指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ret
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea(arr)
	fmt.Println(ret)
}
