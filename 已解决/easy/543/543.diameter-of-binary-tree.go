package problem543

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
        1
       / \
      2   3
     / \
    4   5
*/
type TreeNode = kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := Dfs(root)
	return res

}

func Dfs(root *TreeNode) (length, diameter int) {
	if root == nil {
		return
	}
	leftLen, leftDia := Dfs(root.Left)
	rightLen, rightDia := Dfs(root.Right)

	length = max(leftLen, rightLen) + 1
	diameter = max(leftLen+rightLen, max(leftDia, rightDia))
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
