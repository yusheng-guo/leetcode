package main

import (
	"strconv"
)

// 2409. 统计共同度过的日子数
// https://leetcode.cn/problems/count-days-spent-together/
// 方法二
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	left, right := 0, 0
	if days(arriveAlice) < days(arriveBob) {
		left = days(arriveBob)
	} else {
		left = days(arriveAlice)
	}

	if days(leaveAlice) < days(leaveBob) {
		right = days(leaveAlice)
	} else {
		right = days(leaveBob)
	}
	if right-left+1 > 0 {
		return right - left + 1
	}
	return 0
}

// days 计算从年初开始到date的天数
func days(date string) int {
	daysOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // 每个月的天数
	prefixSum := []int{0}
	for i, v := range daysOfMonth {
		prefixSum = append(prefixSum, v+prefixSum[i])
	}
	idx, _ := strconv.Atoi(date[:2])
	n, _ := strconv.Atoi(date[3:])
	return prefixSum[idx-1] + n
}

// 方法一
//func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
//	aliceArriveM, _ := strconv.Atoi(arriveAlice[:2])
//	aliceArriveD, _ := strconv.Atoi(arriveAlice[3:])
//	aliceLeaveM, _ := strconv.Atoi(leaveAlice[:2])
//	aliceLeaveD, _ := strconv.Atoi(leaveAlice[3:])
//
//	bobArriveM, _ := strconv.Atoi(arriveBob[:2])
//	bobArriveD, _ := strconv.Atoi(arriveBob[3:])
//	bobLeaveM, _ := strconv.Atoi(leaveBob[:2])
//	bobLeaveD, _ := strconv.Atoi(leaveBob[3:])
//
//	leftM, leftD := 0, 0
//	if aliceArriveM == bobArriveM {
//		leftM = aliceArriveM
//		if aliceArriveD >= bobArriveD {
//			leftD = aliceArriveD
//		} else {
//			leftD = bobArriveD
//		}
//	} else if aliceArriveM > bobArriveM {
//		leftM = aliceArriveM
//		leftD = aliceArriveD
//	} else {
//		leftM = bobArriveM
//		leftD = bobArriveD
//	}
//
//	rightM, rightD := 0, 0
//	if aliceLeaveM == bobLeaveM {
//		rightM = aliceLeaveM
//		if aliceLeaveD >= bobLeaveD {
//			rightD = bobLeaveD
//		} else {
//			rightD = aliceLeaveD
//		}
//	} else if aliceLeaveM > bobLeaveM {
//		rightM = bobLeaveM
//		rightD = bobLeaveD
//	} else {
//		rightM = aliceLeaveM
//		rightD = aliceLeaveD
//	}
//
//	fmt.Println(leftM, leftD)
//	fmt.Println(rightM, rightD)
//
//	// 计算相隔天数
//	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // 每个月的天数
//	a, b := leftD, rightD
//	for leftM > 1 {
//		a += days[leftM-2]
//		leftM--
//	}
//	for rightM > 1 {
//		b += days[rightM-2]
//		rightM--
//	}
//	if b-a+1 > 0 {
//		return b - a + 1
//	}
//	return 0
//}

// func main() {
// 	//arriveAlice := "08-15"
// 	//leaveAlice := "08-18"
// 	//arriveBob := "08-16"
// 	//leaveBob := "08-19"
// 	//arriveAlice := "10-01"
// 	//leaveAlice := "10-31"
// 	//arriveBob := "11-01"
// 	//leaveBob := "12-31"
// 	//arriveAlice := "09-01"
// 	//leaveAlice := "10-19"
// 	//arriveBob := "06-19"
// 	//leaveBob := "10-20"
// 	arriveAlice := "03-05"
// 	leaveAlice := "07-14"
// 	arriveBob := "04-14"
// 	leaveBob := "09-21"
// 	ret := countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob)
// 	fmt.Println(ret)
// }
