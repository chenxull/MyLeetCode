package problem110

import "math"

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自底向上,当返回的结果不为-1 时，即可
func isBalanced(root *TreeNode) bool {
	return dep(root) != -1
}

func dep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := dep(root.Left)
	if leftDep == -1 {
		return -1
	}

	rightDep := dep(root.Right)
	if rightDep == -1 {
		return -1
	}

	if math.Abs(float64(rightDep-leftDep)) > 1 {
		return -1
	}
	return max(leftDep, rightDep) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
