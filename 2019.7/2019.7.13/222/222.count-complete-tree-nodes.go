package problem222

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0

	recur(root, &count)
	return count
}
func recur(root *TreeNode, count *int) {
	if root == nil {
		return
	}
	//节点数+1
	*count++

	recur(root.Left, count)
	recur(root.Right, count)

}
