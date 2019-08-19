package problem82

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	dummyhead := &ListNode{Next: head}

	cur := head
	// 用来保存左右去重的节点
	left := dummyhead

	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			left = cur
			cur = cur.Next
		} else {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			left.Next = cur.Next
			cur = cur.Next
		}
	}
	return dummyhead.Next
}
