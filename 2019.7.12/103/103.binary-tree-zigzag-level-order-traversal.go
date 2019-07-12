package problem103

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		// 根据层数来放置节点
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			// 每次插入一个元素，都需要移动整层的数组，将其放置在数组的第一位
			temp := make([]int, len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			//  将移动好的数组放回去
			res[level] = temp
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
