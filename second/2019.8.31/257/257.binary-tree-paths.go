package problem257

import "strconv"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := make([]string, 0, 10)
	var dfs func(root *TreeNode, path string)

	dfs = func(root *TreeNode, path string) {
		if path == "" {
			path = strconv.Itoa(root.Val)
		} else {
			path += "->" + strconv.Itoa(root.Val)
		}

		if root.Left != nil {
			dfs(root.Left, path)
		}

		if root.Right != nil {
			dfs(root.Right, path)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, path)
		}
	}

	dfs(root, "")
	return res
}
