package main

import "fmt"

// TreeNode 定义二叉树结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Queue 定义队列结构
type Queue struct {
	nodes []*TreeNode
}

// Enqueue 入队操作
func (q *Queue) Enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

// Dequeue 出队操作
func (q *Queue) Dequeue() *TreeNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

// BFS 广度优先搜索算法
func BFS(root *TreeNode) {
	if root == nil {
		return // 空树直接返回
	}
	queue := Queue{}    // 创建一个空队列
	queue.Enqueue(root) // 将根节点入队

	for !queue.IsEmpty() { // 当队列不为空时循环执行以下操作：
		node := queue.Dequeue() // 出队一个节点并打印其值
		fmt.Println(node.Val)

		if node.Left != nil { // 如果该节点有左子节点，则将左子节点入队
			queue.Enqueue(node.Left)
		}

		if node.Right != nil { // 如果该节点有右子节点，则将右子节点入队
			queue.Enqueue(node.Right)
		}
	}
}

//func kthLargestLevelSum(root *TreeNode, k int) int64 {
//
//}

func main() {
	// ...
}
