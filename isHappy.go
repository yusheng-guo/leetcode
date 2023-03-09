package main

import "fmt"

// 202. 快乐数
// https://leetcode.cn/problems/happy-number/description/

// 方法四 数学
func isHappy(n int) bool {
	cycle := map[int]bool{4: true, 6: true, 37: true, 58: true, 89: true, 145: true, 42: true, 20: true}
	for n != 1 && !cycle[n] {
		n = step(n)
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 方法三 快慢指针
//func isHappy(n int) bool {
//	slow, fast := n, step(n)
//	for fast != 1 && slow != fast {
//		slow = step(slow)
//		fast = step(step(fast))
//	}
//	return fast == 1
//}
//
//func step(x int) int {
//	end := 0
//	for x != 0 {
//		end += (x % 10) * (x % 10)
//		x /= 10
//	}
//	return end
//}

// 方法二 哈希表优化
//func isHappy(n int) bool {
//	m := make(map[int]bool)
//	for n != 1 && !m[n] {
//		m[n] = true
//		n = step(n)
//	}
//	return n == 1
//}
//
//func step(x int) int {
//	end := 0
//	for x != 0 {
//		end += (x % 10) * (x % 10)
//		x /= 10
//	}
//	return end
//}

// 方法一 哈希表
//func isHappy(n int) bool {
//	m := make(map[int]bool)
//	for n != 1 {
//		n = step(n)
//		if m[n] {
//			return false
//		}
//		m[n] = true
//	}
//	return true
//}
//
//// 每个位置上的数字的平方和
//func step(x int) int {
//	end := 0
//	for x != 0 {
//		end += (x % 10) * (x % 10)
//		x /= 10
//	}
//	return end
//}

func main() {
	fmt.Println(isHappy(19))
}
