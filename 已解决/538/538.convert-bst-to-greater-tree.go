package problem538

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	travel(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	travel(root.Left, sum)

}
