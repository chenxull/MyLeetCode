package problem107

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {

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
	return reverse(res)
}

func reverse(res [][]int) [][]int {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
