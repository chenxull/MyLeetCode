package problem98

import (
	"math"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

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
type TreeNode = kit.TreeNode

// 思路上比较难，之前理解错了题目
func isValidBST(root *TreeNode) bool {

	// golang 中的最小值和最大值
	min, max := math.MinInt64, math.MaxInt64
	return helper(int(min), int(max), root)
}
func helper(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return root.Val > min &&
		root.Val < max &&
		helper(min, root.Val, root.Left) &&
		helper(root.Val, max, root.Right)
}
