package problem102

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	res := [][]int{}

	queue := []*TreeNode{}

	// 将根节点放入队列中
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		temp := []int{}
		for i := 0; i < size; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			temp = append(temp, queue[0].Val)
			queue = queue[1:]
		}
		res = append(res, temp)
	}
	return res

}
