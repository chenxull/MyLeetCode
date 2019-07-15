package problem129

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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

func sumNumbers(root *TreeNode) int {
	res := 0

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}

		temp = temp*10 + root.Val

		if root.Left == nil && root.Right == nil {
			res += temp
			return
		}

		dfs(root.Left, temp)
		dfs(root.Right, temp)
	}

	dfs(root, 0)
	return res
}
