package problem60

import "strconv"

/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	res := strconv.Itoa(t.Val)

	if t.Left == nil && t.Right == nil {
		return res
	}

	res += "(" + tree2str(t.Left) + ")"

	if t.Right != nil {
		res += "(" + tree2str(t.Right) + ")"
	}
	return res
}
