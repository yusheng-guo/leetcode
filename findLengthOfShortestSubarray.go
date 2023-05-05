package main

// 1574. 删除最短的子数组使剩余数组有序
// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
// 双指针 参考官方题解
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	res := j
	for i := 0; i < n; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		if res > j-i-1 {
			res = j - i - 1
		}
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}

// func main() {
// 	arr := []int{1, 2, 3, 10, 4, 2, 3, 5}
// 	ret := findLengthOfShortestSubarray(arr)
// 	fmt.Println(ret)
// }
