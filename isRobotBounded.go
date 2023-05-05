package main

// 1041. 困于环中的机器人
// https://leetcode.cn/problems/robot-bounded-in-circle/
func isRobotBounded(instructions string) bool {
	dir := 90    // 方向 默认向北
	x, y := 0, 0 // 机器人位置
	for _, v := range instructions {
		switch v {
		case 'G': // 移动
			switch dir {
			case 0:
				x++
			case 90:
				y++
			case 180:
				x--
			case 270:
				y--
			}
		case 'L': // 左转
			dir = (dir + 90) % 360
		case 'R': // 右转
			dir = (dir + 270) % 360
		}
	}
	return (x == 0 && y == 0) || dir != 90
}

// func main() {
// 	ret := isRobotBounded("GLGLGGLGL")
// 	fmt.Println(ret)
// }
