package problem148

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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

//  使用归并排序,先对链表进行二分
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 对链表进行切分
	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

func split(head *ListNode) (left, right *ListNode) {
	slowPre := head
	slow, fast := head, head
	//  设置快慢指针，将链表分为二段.fast 来用判断遍历是否结束
	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}
	// 将 left 链表的尾节点的后一个设置为 nil，成为一个独立的链表
	slowPre.Next = nil
	left, right = head, slow
	return

}

func merge(left, right *ListNode) *ListNode {
	cur := &ListNode{}
	headpre := cur

	//  对左右链表进行 merge 操作
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	//  当其中一个链表遍历完成，另一个链表还剩一个元素
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return headpre.Next
}
