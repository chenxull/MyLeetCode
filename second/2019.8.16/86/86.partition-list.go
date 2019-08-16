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

// 这道题目和之间一题很相似，对奇偶位置的元素进行处理，这里条件换成了判断大小
func partition(head *ListNode, x int) *ListNode {
	lesshead := &ListNode{}
	moreHead := &ListNode{}

	less := lesshead
	more := moreHead

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			more.Next = head
			more = more.Next
		}
	}

	more.Next = nil
	less.Next = moreHead.Next

	head = lesshead.Next
	return head
}
