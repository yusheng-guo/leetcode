package main

import (
	"fmt"
	"sort"
)

func splitNum(num int) int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	}) // 排序
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		if i%2 == 0 {
			a = nums[i] + 10*a
		} else {
			b = nums[i] + 10*b
		}
		fmt.Println(a, b)
	}
	return a + b
}

func main() {
	num := 4325
	ret := splitNum(num)
	fmt.Println(ret)
}
