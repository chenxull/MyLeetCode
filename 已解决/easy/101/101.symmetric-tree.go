package problem101

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

func isSymmetric(root *TreeNode) bool {
	return ismmetric(root, root)
}

func ismmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && ismmetric(t1.Left, t2.Right) && ismmetric(t1.Right, t2.Left)
}
