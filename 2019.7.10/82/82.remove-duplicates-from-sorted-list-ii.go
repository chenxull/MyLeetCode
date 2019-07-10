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

//  只要处理的思路清晰，这种题目还是很容易的
func deleteDuplicates(head *ListNode) *ListNode {
	dumyHead := ListNode{Next: head}

	cur := head
	left := &dumyHead
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
	return dumyHead.Next

}
