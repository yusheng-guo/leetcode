package main

import "fmt"

// 6354. K 件物品的最大和
// https://leetcode.cn/problems/k-items-with-the-maximum-sum/
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numZeros+numOnes {
		return numOnes
	}
	return numOnes*2 + numZeros - k
}

func main() {
	numOnes := 3
	numZeros := 3
	numNegOnes := 5
	k := 11
	ret := kItemsWithMaximumSum(numOnes, numZeros, numNegOnes, k)
	fmt.Println(ret)
}
