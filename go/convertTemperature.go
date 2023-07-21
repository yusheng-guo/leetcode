package main

// 2469. 温度转换
// https://leetcode.cn/problems/convert-the-temperature/
func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00
	return []float64{kelvin, fahrenheit}
}

// func main() {
// 	celsius := 36.50
// 	ret := convertTemperature(celsius)
// 	fmt.Println(ret)
// }
