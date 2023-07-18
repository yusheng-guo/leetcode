package main

// https://leetcode.cn/problems/the-employee-that-worked-on-the-longest-task/
// 2432. 处理用时最长的那个任务的员工
func hardestWorker(n int, logs [][]int) int {
	start := 0 // 员工开始任务时刻
	spend := 0 // 处理用时最长时间
	longestTimeConsumingWorkerId := 0
	for _, v := range logs {
		if v[1]-start > spend { // 当前员工完成任务花费的时间 < 之前员工花费的最长时间
			spend = v[1] - start
			longestTimeConsumingWorkerId = v[0]
		}
		if v[1]-start == spend && v[0] < longestTimeConsumingWorkerId {
			longestTimeConsumingWorkerId = v[0]
		}
		start = v[1]
	}
	return longestTimeConsumingWorkerId
}

// func main() {
// 	n := 26
// 	// logs := [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}
// 	logs := [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}
// 	ret := hardestWorker(n, logs)
// 	fmt.Println(ret)
// }
