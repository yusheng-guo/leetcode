package main

// 有效的括号
// https://leetcode.cn/problems/valid-parentheses/

// 方法二 map
func isValid(s string) bool {
	var stack []byte
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 { // 右括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else { // 左括号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 方法一 switch语句
//func isValid(s string) bool {
//	var stack []rune
//	for _, w := range s {
//		switch w {
//		case '(', '[', '{':
//			stack = append(stack, w)
//		case ')':
//			if len(stack) > 0 && stack[len(stack)-1] == '(' {
//				stack = stack[0 : len(stack)-1]
//			} else {
//				return false
//			}
//		case ']':
//			if len(stack) > 0 && stack[len(stack)-1] == '[' {
//				stack = stack[0 : len(stack)-1]
//			} else {
//				return false
//			}
//		case '}':
//			if len(stack) > 0 && stack[len(stack)-1] == '{' {
//				stack = stack[0 : len(stack)-1]
//			} else {
//				return false
//			}
//		}
//	}
//	return len(stack) == 0
//}

// func main() {
// 	input := "()[]{}"
// 	fmt.Println(isValid(input))
// }
