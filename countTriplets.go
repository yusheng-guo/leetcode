package main

import (
	"fmt"
)

// 982. 按位与为零的三元组
// https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/description/

// 方法二 二重枚举 + 哈希表
func countTriplets(nums []int) int {
	l := len(nums)
	cnt := make(map[int]int)
	count := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			cnt[nums[i]&nums[j]]++
		}
	}
	for k := 0; k < l; k++ {
		for x, v := range cnt {
			if x&nums[k] == 0 {
				count += v
			}
		}
	}
	return count
}

// 方法一 暴力枚举
//func countTriplets(nums []int) int {
//	l := len(nums)
//	count := 0
//	for i := 0; i < l; i++ {
//		for j := 0; j < l; j++ {
//			for k := 0; k < l; k++ {
//				if nums[i]&nums[j]&nums[k] == 0 {
//					count++
//				}
//			}
//		}
//	}
//	return count
//}

func main() {
	//nums := []int{29186, 16988, 10334, 34575, 31933, 36792, 49511, 57925, 32412, 20451, 30598, 53021, 57345, 53242, 12320, 65498, 41490, 5470, 47549, 64672, 26333, 44467, 20720, 13993, 47010, 154, 53918, 34625, 6408, 7792, 42550, 57255, 18348, 64127, 6565, 55963, 25997, 5331, 18003, 29686, 51926, 56450, 8743, 34135, 42649, 22744, 57138, 18245, 1217, 2664, 44629, 36592, 62029, 61501, 59628, 2752, 49095, 45962, 42949, 46685, 59160, 41952, 24408, 35253, 13621, 38477, 583, 9897, 21850, 51748, 47087, 22167, 40667, 2700, 10738, 64350, 18974, 25873, 35945, 19228, 31330, 15709, 3162, 11769, 34278, 4644, 26422, 21529, 27179, 6514, 59549, 31064, 55105, 63978, 22521, 31397, 45848, 27953, 19785, 48152}
	nums := []int{2, 1, 3}
	ret := countTriplets(nums)
	fmt.Println(ret)
}
