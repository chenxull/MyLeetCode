package problem98

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	min, max := math.MinInt64, math.MaxInt64

	return helper(min, max, root)
}

func helper(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return root.Val > min && root.Val < max &&
		helper(min, root.Val, root.Left) &&
		helper(root.Val, max, root.Right)
}
