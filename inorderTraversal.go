package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 方法二 迭代

// 方法一 递归
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	var ret []int
//	ret = append(ret, inorderTraversal(root.Left)...)
//	ret = append(ret, root.Val)
//	ret = append(ret, inorderTraversal(root.Right)...)
//	return ret
//}
