package main

// 方法一 将
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	for i := 0; i < n; i++ {
//		nums1[m+i] = nums2[i]
//	}
//	sort.Slice(nums1, func(i, j int) bool {
//		if nums1[i] < nums1[j] {
//			return true
//		}
//		return false
//	})
//}

// 方法二 双指针
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	p1, p2 := 0, 0
//	sorted := make([]int, 0, m+n)
//	for {
//		if p1 == m {
//			sorted = append(sorted, nums2[p2:]...)
//			break
//		}
//		if p2 == n {
//			sorted = append(sorted, nums1[p1:]...)
//			break
//		}
//		if nums1[p1] < nums2[p2] {
//			sorted = append(sorted, nums1[p1])
//			p1++
//		} else {
//			sorted = append(sorted, nums2[p2])
//			p2++
//		}
//	}
//	copy(nums1, sorted)
//}

// 方法三：逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, cur := m-1, n-1, m+n-1
	for {
		if p1 == -1 {
			copy(nums1, nums2[p2:])
			break
		}
		if p2 == -1 {
			break
		}
		if nums1[p1] < nums2[p2] {
			nums1[cur] = nums2[p2]
			cur--
			p2--
		} else {
			nums1[cur] = nums1[p1]
			cur--
			p1--
		}
	}
}

//
//func main() {
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//	merge(nums1, len(nums1), nums2, len(nums2))
//}
