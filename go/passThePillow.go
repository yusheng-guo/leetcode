package main

const (
	Forward  = true  // 向前
	Backward = false // 向后
)

// 6307. 递枕头
func passThePillow(n int, time int) int {
	direction := Forward // 默认下一个传递
	position := 1        // 默认枕头在第一个人手里
	for time != 0 {
		if direction == Forward {
			position++
		}
		if direction == Backward {
			position--
		}
		if position == 1 {
			direction = Forward
		}
		if position == n {
			direction = Backward
		}
		time--
		//fmt.Println(position, direction, time)
	}
	return position
}

// func main() {
// 	fmt.Println(passThePillow(4, 5))
// }
