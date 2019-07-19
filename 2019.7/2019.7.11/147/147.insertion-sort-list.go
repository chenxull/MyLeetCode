package problem147

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
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

func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		//  位置不用动
		if cur.Val < p.Val {
			cur = p
			continue
		}

		//  保留后续节点的指针
		cur.Next = p.Next
		//  需要在 headPre 与 cur 之间找到一个合适的位置
		pre, next := headPre, headPre.Next
		//  符合条件的 pre.val < p.val <= next.val
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}

	return headPre.Next
}
