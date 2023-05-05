package main

// 6361. 对角线上的质数
// https://leetcode.cn/contest/weekly-contest-340/problems/prime-in-diagonal/
func diagonalPrime(nums [][]int) int {
	ans := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if isPrime(nums[i][i]) && nums[i][i] > ans {
			ans = nums[i][i]
		}

		if isPrime(nums[i][l-i-1]) && nums[i][l-i-1] > ans {
			ans = nums[i][l-i-1]
		}
	}
	return ans
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	arr := [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}
// 	ret := diagonalPrime(arr)
// 	fmt.Println(ret)
// }
