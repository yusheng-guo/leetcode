package main

import "fmt"

//func ballGame(num int, plate []string) [][]int {
//	var ret [][]int
//	lx, ly := len(plate[0]), len(plate)
//	// 上边
//	for i := 1; i < lx-1; i++ {
//		// 确定方向
//		direction := 90 // 正右为0 顺时针加 逆时针减
//		switch plate[0][i] {
//		case 'W':
//			direction = (direction + 270) % 360
//		case 'E':
//			direction = (direction + 90) % 360
//		}
//		// 上边缘
//		if roll(i, 0, lx, ly, direction, num, plate) {
//			ret = append(ret, []int{0, i})
//		}
//	}
//	// 下边
//	for i := 1; i < lx-1; i++ {
//		// 确定方向
//		direction := 270 // 正右为0 顺时针加 逆时针减
//		switch plate[ly-1][i] {
//		case 'W':
//			direction = (direction + 270) % 360
//		case 'E':
//			direction = (direction + 90) % 360
//		}
//		// 下边缘
//		if roll(i, ly-1, lx, ly, direction, num, plate) {
//			ret = append(ret, []int{ly - 1, i})
//		}
//	}
//	// 左边
//	for i := 1; i < ly-1; i++ {
//		// 确定方向
//		direction := 0 // 正右为0 顺时针加 逆时针减
//		switch plate[i][0] {
//		case 'W':
//			direction = (direction + 270) % 360
//		case 'E':
//			direction = (direction + 90) % 360
//		}
//		if roll(0, i, lx, ly, direction, num, plate) {
//			ret = append(ret, []int{i, 0})
//		}
//	}
//	// 右边
//	for i := 1; i < ly-1; i++ {
//		// 确定方向
//		direction := 180 // 正右为0 顺时针加 逆时针减
//		switch plate[i][lx-1] {
//		case 'W':
//			direction = (direction + 270) % 360
//		case 'E':
//			direction = (direction + 90) % 360
//		}
//		// 下边缘
//		if roll(lx-1, i, lx, ly, direction, num, plate) {
//			ret = append(ret, []int{i, lx - 1})
//		}
//	}
//	return ret
//}
//
//func roll(x, y, lx, ly, direction, pace int, plate []string) bool {
//	for pace > 0 {
//		// 移动
//		switch direction {
//		case 0:
//			x++
//		case 90:
//			y++
//		case 180:
//			x--
//		case 270:
//			y--
//		}
//		if x < 0 || x > lx-1 || y < 0 || y > ly-1 {
//			break
//		}
//		// 方向
//		switch plate[y][x] {
//		case 'W':
//			direction = (direction + 270) % 360
//		case 'E':
//			direction = (direction + 90) % 360
//		case '.':
//			if x == lx-1 || y == ly-1 {
//				break
//			}
//		case 'O':
//			return true
//		}
//		pace--
//	}
//	return false
//}

func ballGame(num int, plate []string) [][]int {
	var ret [][]int
	lx, ly := len(plate[0]), len(plate) // 横向长度 纵向长度
	// 上边
	for i := 1; i < lx-1; i++ {
		if roll(i, 0, lx, ly, 90, num, plate) {
			ret = append(ret, []int{0, i})
		}
	}
	// 下边
	for i := 1; i < lx-1; i++ {
		if roll(i, ly-1, lx, ly, 270, num, plate) {
			ret = append(ret, []int{ly - 1, i})
		}
	}
	// 左边
	for i := 1; i < ly-1; i++ {
		if roll(0, i, lx, ly, 0, num, plate) {
			ret = append(ret, []int{i, 0})
		}
	}
	// 右边
	for i := 1; i < ly-1; i++ {
		if roll(lx-1, i, lx, ly, 180, num, plate) {
			ret = append(ret, []int{i, lx - 1})
		}
	}
	return ret
}

// roll 从(x,y)开始在plate上滚动 能在pace步之内进洞 返回true 否则返回false
func roll(x, y, lx, ly, direction, pace int, plate []string) bool {
	if plate[y][x] != '.' { // 从非空白区域开始 直接返回false
		return false
	}
	for pace > 0 {
		// 移动
		switch direction {
		case 0:
			x++
		case 90:
			y++
		case 180:
			x--
		case 270:
			y--
		}
		if x < 0 || x > lx-1 || y < 0 || y > ly-1 {
			break
		}
		// 方向
		switch plate[y][x] {
		case 'W':
			direction = (direction + 270) % 360
		case 'E':
			direction = (direction + 90) % 360
		case '.':
			if x == lx-1 || y == ly-1 {
				break
			}
		case 'O':
			return true
		}
		pace--
	}
	return false
}

func main() {
	num := 41
	plate := []string{"E...W..WW", ".E...O...", "...WO...W", "..OWW.O..", ".W.WO.W.E", "O..O.W...", ".OO...W..", "..EW.WEE."}
	ret := ballGame(num, plate)
	for i := 0; i < len(plate); i++ {
		for j := 0; j < len(plate[0]); j++ {
			fmt.Printf("%c ", plate[i][j])
		}
		fmt.Println()
	}
	fmt.Println(ret)
}
