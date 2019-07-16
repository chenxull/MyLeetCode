package problem236

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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

	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// 当l，r 同时不为 nil 时，说明肯定l，r 中肯定遇到了 p 和q，所以才导致这二个
	// 参数不为空，这节返回当前节点即可

	if l != nil && r != nil {
		return root
	}

	//  如果左子树中没有找到相应的p 或 q，说明 p 和 q 在 root.right中
	if l == nil {
		return r
	}

	return l

}
