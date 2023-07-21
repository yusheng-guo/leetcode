package main

// 2574. 左右元素和的差值
// https://leetcode.cn/problems/left-and-right-sum-differences/
func leftRigthDifference(nums []int) []int {
	ans := make([]int, len(nums))
	sum := 0 // 计算总和
	for _, v := range nums {
		sum += v
	}
	sumLeft, sumRight := 0, sum
	for i := 0; i < len(nums); i++ {
		sumRight -= nums[i]
		diff := sumRight - sumLeft
		if diff < 0 { // 取绝对值
			diff = -diff
		}
		ans[i] = diff
		sumLeft += nums[i]
	}
	return ans
}

// 方法一
//func leftRigthDifference(nums []int) []int {
//	var ans []int
//	sum := 0 // 计算总和
//	for _, v := range nums {
//		sum += v
//	}
//	sumLeft, sumRight := 0, sum
//	for _, v := range nums {
//		sumRight -= v
//		diff := sumRight - sumLeft
//		if diff < 0 { // 取绝对值
//			diff = -diff
//		}
//		ans = append(ans, diff)
//		sumLeft += v
//	}
//	return ans
//}

// func main() {
// 	nums := []int{10, 4, 8, 3}
// 	ret := leftRigthDifference(nums)
// 	fmt.Println(ret)
// }
