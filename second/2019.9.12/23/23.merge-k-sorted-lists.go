package problem23

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	length := len(lists)
	mid := length / 2

	if length == 1 {
		return lists[0]
	}

	if length == 2 {
		l0, l1 := lists[0], lists[1]
		var res, cur *ListNode

		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}

		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}

		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next, l0 = l0, l0.Next
			} else {
				cur.Next, l1 = l1, l1.Next
			}
			cur = cur.Next
		}

		if l0 == nil {
			cur.Next = l1
		}

		if l1 == nil {
			cur.Next = l0
		}
		return res

	}

	return mergeKLists([]*ListNode{mergeKLists(lists[mid:]), mergeKLists(lists[:mid])})

}
