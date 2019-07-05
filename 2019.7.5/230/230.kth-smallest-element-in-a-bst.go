package problem230

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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

var number int
var count int

func kthSmallest(root *TreeNode, k int) int {
	count = k
	helper(root)
	return number
}

func helper(root *TreeNode) {
	if root.Left != nil {
		helper(root.Left)
	}
	count--
	if count == 0 {
		number = root.Val
		return
	}
	if root.Right != nil {
		helper(root.Right)
	}
}
