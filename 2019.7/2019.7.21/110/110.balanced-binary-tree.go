package problem110

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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

func isBalanced(root *TreeNode) bool {
	_, isbalanced := recur(root)
	return isbalanced
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := recur(root.Left)
	rightHeight, rightBalanced := recur(root.Right)

	if leftBalanced && rightBalanced && abs(leftHeight, rightHeight) <= 1 {
		return max(leftHeight, rightHeight) + 1, true
	}

	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
