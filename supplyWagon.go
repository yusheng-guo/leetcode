package main

import "fmt"

// 1. 补给马车

// 方法二 动态规划

// 方法一 枚举
func supplyWagon(supplies []int) []int {
	half := len(supplies) / 2
	for len(supplies) > half {
		adjacentSum := supplies[0] + supplies[1] // 相邻和
		index := 0                               // 相邻和最小的索引
		for i := 1; i < len(supplies)-1; i++ {
			if supplies[i]+supplies[i+1] < adjacentSum {
				adjacentSum = supplies[i] + supplies[i+1]
				index = i
			}
		}
		supplies[index] = adjacentSum
		if index == len(supplies)-2 { // 删除切片(index+1)值
			supplies = supplies[:index+1]
		} else {
			supplies = append(supplies[:index+1], supplies[index+2:]...)
		}
		fmt.Println(supplies)
	}
	return supplies
}

func main() {
	input := []int{1, 3}
	fmt.Println(supplyWagon(input))
}
