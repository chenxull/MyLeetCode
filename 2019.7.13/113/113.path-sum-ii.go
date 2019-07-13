package problem113

// import (
// 	"github.com/aQuaYi/LeetCode-in-Go/kit"
// )

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
//  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level int, sum int) {
		if node == nil {
			return
		}
		// 开始阶段path 的长度为 0，需要根据层次信息，来将节点放置在合适的位置
		if level >= len(path) {
			//  error，无效内存空间
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}

		sum -= node.Val

		if sum == 0 && node.Left == nil && node.Right == nil {
			temp := make([]int, level+1)
			// //  为了避免二叉树不是满的情况，会把一些不属于 path 的结果给放入到 res 中
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)

	}
	dfs(root, 0, sum)

	return res
}
