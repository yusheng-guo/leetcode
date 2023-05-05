package main

// LCP 61. 气温变化趋势
// https://leetcode.cn/problems/6CE719/?envType=study-plan&id=2022-qiu-sai-ti-ji&plan=lccup&plan_progress=x5979eah
//const (
//	up = iota
//	intact
//	down
//)

// 方法二 简化
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	l := len(temperatureA)
	ret := 0
	var rets []int
	for i := 1; i < l; i++ {
		if (temperatureA[i] > temperatureA[i-1] && temperatureB[i] > temperatureB[i-1]) ||
			(temperatureA[i] == temperatureA[i-1] && temperatureB[i] == temperatureB[i-1]) ||
			(temperatureA[i] < temperatureA[i-1] && temperatureB[i] < temperatureB[i-1]) {
			ret++
			continue
		}
		rets = append(rets, ret)
		ret = 0
	}
	for i := 0; i < len(rets); i++ {
		if ret < rets[i] {
			ret = rets[i]
		}
	}
	return ret
}

// 方法一
//func temperatureTrend(temperatureA []int, temperatureB []int) int {
//	l := len(temperatureA)
//	ret := 0
//	var rets []int
//	for i := 1; i < l; i++ {
//		if temperatureA[i] > temperatureA[i-1] && temperatureB[i] > temperatureB[i-1] {
//			ret++
//			continue
//		}
//		if temperatureA[i] == temperatureA[i-1] && temperatureB[i] == temperatureB[i-1] {
//			ret++
//			continue
//		}
//		if temperatureA[i] < temperatureA[i-1] && temperatureB[i] < temperatureB[i-1] {
//			ret++
//			continue
//		}
//		rets = append(rets, ret)
//		ret = 0
//	}
//	for i := 0; i < len(rets); i++ {
//		if ret < rets[i] {
//			ret = rets[i]
//		}
//	}
//	return ret
//}

// func main() {
// 	temperatureA := []int{21, 18, 18, 18, 31}
// 	temperatureB := []int{34, 32, 16, 16, 17}
// 	ret := temperatureTrend(temperatureA, temperatureB)
// 	fmt.Println(ret)
// }
