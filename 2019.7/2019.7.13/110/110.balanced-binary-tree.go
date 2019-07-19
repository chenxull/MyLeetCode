package problem

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
//  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//  有二种方式自下而上，和自上而下
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftdep := dep(root.Left)
	rightdp := dep(root.Right)

	return math.Abs(float64(leftdep-rightdp)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func dep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(dep(root.Left), dep(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
