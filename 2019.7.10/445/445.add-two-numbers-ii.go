package problem445

/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 将数组放入链表中，来进行处理
	s1 := make([]int, 0, 128)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	s2 := make([]int, 0, 128)
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	head := &ListNode{}
	sum := 0
	carry := 0
	for len(s1) != 0 || len(s2) != 0 {
		sum = carry

		if len(s1) != 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) != 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		carry = sum / 10

		head.Val = sum % 10
		//  当二个数组为[5] [5] 时，确保可以先将进位信息保存下来。
		front := &ListNode{Val: sum / 10}
		front.Next = head
		head = front
	}
	if head.Val == 0 {
		return head.Next
	}
	return head
}
