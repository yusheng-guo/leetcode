package main

import "fmt"

// 1615. 最大网络秩
// https://leetcode.cn/problems/maximal-network-rank/
// 如果 i 与 j 之间没有道路连接，则此时 (i,j) 的网络秩为 degree[i]+degree[j]
// 如果 iii 与 jjj 之间存在道路连接，则此时 (i,j)(i,j)(i,j) 的网络秩为 degree[i]+degree[j]−1

// 方法一 枚举
func maximalNetworkRank(n int, roads [][]int) int {
	connect := make([][]int, n) // 两座城市之间是否有通路
	for i := range connect {
		connect[i] = make([]int, n)
	}
	degree := make([]int, n) // 连接某个城市的城市个数 (度）
	for _, v := range roads {
		connect[v[0]][v[1]] = 1
		connect[v[1]][v[0]] = 1
		degree[v[0]] += 1
		degree[v[1]] += 1
	}
	fmt.Println(connect)
	fmt.Println(degree)
	maxRank := 0 // 最大秩
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degree[i] + degree[j] - connect[i][j]
			if rank > maxRank {
				maxRank = rank
			}
		}
	}
	return maxRank
}

func main() {
	n := 5
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}
	ret := maximalNetworkRank(n, roads)
	fmt.Println(ret)
}
