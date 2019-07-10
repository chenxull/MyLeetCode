package problem86

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
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

//   创建二个新的链表来分别保存数据
func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{}
	afterHead := &ListNode{}

	before := beforeHead
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next

	head = beforeHead.Next
	return head
}
