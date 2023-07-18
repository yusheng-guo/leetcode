package main

// 1287. 有序数组中出现次数超过25%的元素
// https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/
// 方法二 二叉查找
func findSpecialInteger(arr []int) int {
	l := len(arr)
	span := l/4 + 1
	for i := 0; i < l; i += span {
		curr := arr[i]
		left := binarySearchFirst(arr, curr)
		right := binarySearchLast(arr, curr)
		if right-left+1 >= span {
			return curr
		}
	}
	return -1
}

func binarySearchFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	firstPos := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			firstPos = mid
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstPos
}

func binarySearchLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	lastPos := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			lastPos = mid
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return lastPos
}

// 方法一 遍历
//func findSpecialInteger(arr []int) int {
//	l := len(arr)
//	count := 1
//	for i := 0; i < l-1; i++ {
//		if arr[i] == arr[i+1] {
//			count++
//		} else {
//			count = 1
//		}
//		if count*4 > l {
//			return arr[i]
//		}
//	}
//	return arr[0]
//}

// func main() {
// 	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
// 	ret := findSpecialInteger(arr)
// 	fmt.Println(ret)
// }
