package problem111

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//  需要注意和 104 题不同之处。 叶子节点的定义为，没有孩子的节点。
func minDepth(root *TreeNode) int {

	switch {
	case root == nil:
		return 0

	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
