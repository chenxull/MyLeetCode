package problem160

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	pa, pb := headA, headB

	countA, countB := 0, 0

	for pa != nil {
		pa = pa.Next
		countA++
	}
	for pb != nil {
		pb = pb.Next
		countB++
	}

	if countA <= countB {
		dis := countB - countA
		pa = headA
		pb = headB
		for dis > 0 {
			dis--
			pb = pb.Next
		}
	} else {
		dis := countA - countB
		pa = headA
		pb = headB
		for dis > 0 {
			dis--
			pa = pa.Next
		}
	}

	for pa != pb {
		pa = pa.Next
		pb = pb.Next
	}
	return pa
}
