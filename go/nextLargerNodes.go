package main

// 1019. 链表中的下一个更大节点
// https://leetcode.cn/problems/next-greater-node-in-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	var ans []int
	for head != nil {
		curr := head
		for curr.Next != nil && curr.Next.Val <= head.Val {
			curr = curr.Next
		}
		if curr.Next == nil {
			ans = append(ans, 0)
		} else {
			ans = append(ans, curr.Next.Val)
		}
		head = head.Next
	}
	return ans
}

// func main() {
// 	//nextLargerNodes
// }
