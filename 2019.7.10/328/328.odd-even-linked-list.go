package problem328

/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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

func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{}
	evenHead := &ListNode{}

	odd := oddHead
	even := evenHead
	count := 1
	for head != nil {
		if count%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		count++
		head = head.Next
	}

	even.Next = nil
	odd.Next = evenHead.Next

	head = oddHead.Next
	return head
}
