package problem237

/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//  因为题目给的条件有限 ，所以不能使用常规的方法来求解。
//  直接使用赋值的方法
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	if node.Next == nil {
		node = nil
		return
	}

	delNode := node.Next
	node.Val = delNode.Val

	node.Next = delNode.Next
}
