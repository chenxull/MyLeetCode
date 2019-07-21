package problem111

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := minDepth(root.Left)
	rightHeight := minDepth(root.Right)

	// 为何要这一步？因为要去的最小深度为 root 到叶子节点的长度
	// 如果一个节点只有左子树或右子树，返回其中高度大的。
	if leftHeight == 0 || rightHeight == 0 {
		return leftHeight + rightHeight + 1
	}

	return min(leftHeight, rightHeight) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
