package main

// 1053. 交换一次的先前排列
// https://leetcode.cn/problems/previous-permutation-with-one-swap/
// 贪心算法
func prevPermOpt1(arr []int) []int {
	// 找i
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			// 找j
			j := len(arr) - 1
			for arr[j] >= arr[i] || arr[j] == arr[j-1] {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}

// func main() {
// 	arr := []int{3, 1, 1, 3}
// 	ret := prevPermOpt1(arr)
// 	fmt.Println(ret)
// }
