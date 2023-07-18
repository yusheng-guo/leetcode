package main

import (
	"fmt"
)

func maxNumOfMarkedIndices(nums []int) int {
	l := len(nums)
	index := make([]int, l)
	for i := 0; i < l && index[i] == 0; i++ {
		for j := i + 1; j < l && index[j] == 0; j++ {
			fmt.Println(i, nums[i], j, nums[j])
			if 2*nums[i] <= nums[j] || 2*nums[j] <= nums[i] {
				index[i], index[j] = 1, 1
			}
		}
	}
	ret := 0
	for _, v := range index {
		if v == 1 {
			ret++
		}
	}
	return ret
}

// func main() {
// 	nums := []int{9, 2, 5, 4}
// 	ret := maxNumOfMarkedIndices(nums)
// 	fmt.Println(ret)
// }
