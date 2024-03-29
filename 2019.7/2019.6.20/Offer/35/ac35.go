package offer35

type complexListNode struct {
	value   int
	next    *complexListNode
	sibling *complexListNode
}

// 复制原始链表，并将其中的每个元素连接到原链表对应元素的后面
func cloneNodes(head *complexListNode) {
	pNode := head
	for pNode != nil {
		pCloned := new(complexListNode)
		pCloned.value = pNode.value
		pCloned.next = pNode.next
		// 因为 pCloned 已经保存了 pNode 后续节点的地址
		pNode.next = pCloned
		pNode = pCloned.next
	}
}

// 直接在综合的链表上 处理 sibling 指针
func connectSiblingNodes(head *complexListNode) {
	pNode := head
	for pNode != nil {
		pCloned := pNode.next
		if pNode.sibling != nil {
			pCloned.sibling = pNode.sibling.next
		}
		pNode = pCloned.next
	}
}

// 将大链表拆分为二个链表.
func reconnectNodes(head *complexListNode) *complexListNode {
	pNode := head
	var pClonedHead, pClonedNode *complexListNode
	// 处理好头节点
	if pNode != nil {
		pClonedHead = pNode.next
		pClonedNode = pClonedHead

		pNode.next = pClonedNode.next
		pNode = pNode.next
	}

	for pNode != nil && pClonedNode != nil {
		pClonedNode.next = pNode.next
		pClonedNode = pClonedNode.next

		pNode.next = pClonedNode.next
		pNode = pNode.next
	}
	return pClonedHead

}

func Clone(Head *complexListNode) *complexListNode {
	cloneNodes(Head)
	connectSiblingNodes(Head)
	return reconnectNodes(Head)
}
