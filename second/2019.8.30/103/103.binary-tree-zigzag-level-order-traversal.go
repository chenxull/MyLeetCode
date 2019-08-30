package problem103

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, level int)
	// 前序遍历
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, []int{})
		}

		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			// 奇数层，要从右向左添加
			tmp := make([]int, len(res[level])+1)
			copy(tmp[1:], res[level])
			tmp[0] = root.Val
			res[level] = tmp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
