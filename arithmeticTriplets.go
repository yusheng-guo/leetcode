package main

// 2367. 算术三元组的数目
// https://leetcode.cn/problems/number-of-arithmetic-triplets/description/
// 方法三 三指针
func arithmeticTriplets(nums []int, diff int) int {
	l := len(nums)
	count := 0
	for i, j, k := 0, 1, 2; i < l-2; i++ {
		for j < l-1 && nums[j]-nums[i] < diff { // 定位j
			j++
		}
		for k < l && nums[k]-nums[i] < diff*2 { // 定位k
			k++
		}
		if j > l-2 || k > l-1 {
			break
		}
		if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
			j++
			k++
			count++
		}
	}
	return count
}

// 方法二 哈希表
//func arithmeticTriplets(nums []int, diff int) int {
//	l := len(nums)
//	count := 0
//	m := map[int]struct{}{}
//	for i := 0; i < l; i++ {
//		m[nums[i]] = struct{}{}
//	}
//	for i := l - 1; i >= 0; i-- {
//		_, ok1 := m[nums[i]-diff]
//		_, ok2 := m[nums[i]-diff*2]
//		if ok1 && ok2 {
//			count++
//		}
//
//	}
//	return count
//}

// 方法一 三层for循环
//func arithmeticTriplets(nums []int, diff int) int {
//	l := len(nums)
//	count := 0
//	for i := 0; i < l-2; i++ {
//		for j := i + 1; j < l-1; j++ {
//			for k := i + 2; k < l; k++ {
//				if nums[k]-nums[j] == diff && nums[j]-nums[i] == diff {
//					count++
//				}
//			}
//		}
//	}
//	return count
//}

// func main() {
// 	nums := []int{0, 1, 4, 6, 7, 10}
// 	diff := 3
// 	ret := arithmeticTriplets(nums, diff)
// 	fmt.Println(ret)
// }
