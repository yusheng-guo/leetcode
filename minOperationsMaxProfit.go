package main

import (
	"fmt"
)

// 1599. 经营摩天轮的最大利润
// https://leetcode.cn/problems/maximum-profit-of-operating-a-centennial-wheel/
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	numberOfCustomers := 0 // 上摩天轮的游客数
	waitingCustomers := 0  // 等待摩天轮的游客数
	maxProfit := 0         // 最大利润
	profit := 0            // 当前获取利润
	count := -1            // 最大化利润所需执行的最小轮转次数
	i := 0                 // customers索引
	runningCount := 1      // for循环次数
	for waitingCustomers != 0 || i < len(customers) {
		if i < len(customers) {
			waitingCustomers += customers[i] // 游客抵达
			i++
		}
		if waitingCustomers >= 4 {
			numberOfCustomers += 4
			profit = numberOfCustomers*boardingCost - runningCount*runningCost
			waitingCustomers -= 4 // 摩天轮上四人
		} else {
			numberOfCustomers += waitingCustomers
			profit = numberOfCustomers*boardingCost - runningCount*runningCost
			waitingCustomers = 0
		}
		if profit > maxProfit {
			maxProfit = profit
			count = runningCount
		}
		runningCount++
	}
	return count
}

func main() {
	customers := []int{2}
	boardingCost := 2
	runningCost := 4
	count := minOperationsMaxProfit(customers, boardingCost, runningCost)
	fmt.Println(count)
}
