package problem235

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, p, q int) *TreeNode {
	temp := root.Val
	if temp < p && temp < q {
		//  注意每次递归都需要返回
		return helper(root.Right, p, q)
	} else if temp > p && temp > q {
		return helper(root.Left, p, q)
	}
	return root
}
