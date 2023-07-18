package main

import (
	"strconv"
)

// 682. 棒球比赛
// https://leetcode.cn/problems/baseball-game/description/
// 方法二 栈
func calPoints(operations []string) int {
	sum := 0
	var scores []int // 栈
	for i := 0; i < len(operations); i++ {
		lenScores := len(scores)
		switch operations[i] {
		case "+":
			scores = append(scores, scores[lenScores-1]+scores[lenScores-2])
			sum += scores[lenScores]
		case "D":
			scores = append(scores, scores[lenScores-1]*2)
			sum += scores[lenScores]
		case "C":
			sum -= scores[lenScores-1]
			scores = scores[0 : lenScores-1] // 出栈
		default:
			n, _ := strconv.Atoi(operations[i])
			scores = append(scores, n) // 入栈
			sum += n
		}
	}
	return sum
}

// 方法一 数组
//func calPoints(operations []string) int {
//	var scores []int
//	for i := 0; i < len(operations); i++ {
//		switch operations[i] {
//		case "+":
//			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
//		case "D":
//			scores = append(scores, scores[len(scores)-1]*2)
//		case "C":
//			scores = scores[0 : len(scores)-1]
//		default:
//			score, _ := strconv.Atoi(operations[i])
//			scores = append(scores, score)
//		}
//	}
//	ret := 0
//	for _, v := range scores {
//		ret += v
//	}
//	return ret
//}

// func main() {
// 	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
// 	ret := calPoints(ops)
// 	fmt.Println(ret)
// }
