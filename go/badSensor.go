package main

// 1826. 有缺陷的传感器
// https://leetcode.cn/problems/faulty-sensor/
func badSensor(sensor1 []int, sensor2 []int) int {
	i, j := 0, 0
	for i < len(sensor1) && sensor1[i] == sensor2[j] {
		i++
		j++
	}
	if i >= len(sensor1)-1 { // 都没有缺陷或无法判断
		return -1
	}
	flag1 := false // 标记1
	flag2 := false // 标记2
	origini, originj := i, j
	if sensor1[i+1] == sensor2[j] {
		i++
		for i < len(sensor1) && sensor1[i] == sensor2[j] {
			i++
			j++
		}
		if i == len(sensor1) && j == len(sensor2)-1 {
			flag1 = true
		}
	}
	i, j = origini, originj
	if sensor1[i] == sensor2[j+1] {
		j++
		for j < len(sensor2) && sensor1[i] == sensor2[j] {
			i++
			j++
		}
		if i == len(sensor1)-1 && j == len(sensor2) {
			flag2 = true
		}
	}
	if flag1 && !flag2 {
		return 2
	}
	if !flag1 && flag2 {
		return 1
	}
	return -1
}

// func main() {
// 	sensor1 := []int{2, 3, 2, 2, 3, 2}
// 	sensor2 := []int{2, 3, 2, 3, 2, 7}
// 	ret := badSensor(sensor1, sensor2)
// 	fmt.Println(ret)
// }
