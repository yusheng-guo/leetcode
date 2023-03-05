package main

import (
	"fmt"
	"math/big"
)

// 6309. 分割数组使乘积互质

// 方法三 大数 big库
// 超出时间限制
func findValidSplit(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		product1, product2 := big.NewInt(1), big.NewInt(1)
		for x := 0; x < i+1; x++ {
			product1.Mul(product1, big.NewInt(int64(nums[x])))
		}
		for y := i + 1; y < len(nums); y++ {
			product2.Mul(product2, big.NewInt(int64(nums[y])))
		}
		g := new(big.Int)
		g.GCD(nil, nil, product1, product2)
		if g.Cmp(big.NewInt(1)) == 0 {
			return i
		}
	}
	return -1
}

// 方法二 优化
// 超出时间限制
//func gcd(a, b int) int {
//	if b == 0 { // 如果b为0，返回a
//		return a
//	}
//	return gcd(b, a%b) // 否则递归调用gcd函数，用b和a对b取余作为参数
//}
//
//func findValidSplit(nums []int) int {
//	// 元素乘积
//	for i := 0; i < len(nums)-1; i++ {
//		var product1, product2 []int
//		// 前i+1个元素乘积
//		for x := 0; x < i+1; x++ {
//			product1 = append(product1, nums[x])
//		}
//		for y := i + 1; y < len(nums); y++ {
//			product2 = append(product2, nums[y])
//		}
//		fmt.Println(i, product1, product2)
//		// 判断数组product1乘积和product2乘积是否互质
//		//if gcd(product1, product2) == 1 {
//		//	return i
//		//}
//		isBreak := false
//		for _, v1 := range product1 {
//			for _, v2 := range product2 {
//				if gcd(v1, v2) != 1 {
//					isBreak = true
//					break
//				}
//			}
//			if isBreak {
//				break
//			}
//		}
//		if !isBreak {
//			return len(product1) - 1
//		}
//	}
//	return -1
//}

// 方法一
// 缺点：乘积太大
// gcd 最大公约数
//func gcd(a, b int) int {
//	if b == 0 { // 如果b为0，返回a
//		return a
//	}
//	return gcd(b, a%b) // 否则递归调用gcd函数，用b和a对b取余作为参数
//}
//
//func findValidSplit(nums []int) int {
//	// 元素乘积
//	for i := 0; i < len(nums)-1; i++ {
//		product1, product2 := 1, 1
//		// 前i+1个元素乘积
//		for x := 0; x < i+1; x++ {
//			product1 *= nums[x]
//		}
//		for y := i + 1; y < len(nums); y++ {
//			product2 *= nums[y]
//		}
//		fmt.Println(i, product1, product2)
//		if gcd(product1, product2) == 1 {
//			return i
//		}
//	}
//	return -1
//}

func main() {
	fmt.Println(findValidSplit([]int{557663, 280817, 472963, 156253, 273349, 884803, 756869, 763183, 557663, 964357, 821411, 887849, 891133, 453379, 464279, 574373, 852749, 15031, 156253, 360169, 526159, 410203, 6101, 954851, 860599, 802573, 971693, 279173, 134243, 187367, 896953, 665011, 277747, 439441, 225683, 555143, 496303, 290317, 652033, 713311, 230107, 770047, 308323, 319607, 772907, 627217, 119311, 922463, 119311, 641131, 922463, 404773, 728417, 948281, 612373, 857707, 990589, 12739, 9127, 857963, 53113, 956003, 363397, 768613, 47981, 466027, 981569, 41597, 87149, 55021, 600883, 111953, 119083, 471871, 125641, 922463, 562577, 269069, 806999, 981073, 857707, 831587, 149351, 996461, 432457, 10903, 921091, 119083, 72671, 843289, 567323, 317003, 802129, 612373}))
}
