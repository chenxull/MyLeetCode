package offer36

type binaryTreeNode struct {
	value int
	left  *binaryTreeNode
	right *binaryTreeNode
}

// 本题的关键在于理解 已经转换好的链表中最后一个节点是当前值最大的节点 也就是 lastNodeInList
// 这个连接前后二个节点的关键元素
func Convert(root *binaryTreeNode) *binaryTreeNode {
	var lastNodeInList *binaryTreeNode
	ConvertNode(root, &lastNodeInList)

	headOfConvertList := lastNodeInList
	// 寻找链表的头节点
	for headOfConvertList != nil && headOfConvertList.left != nil {
		headOfConvertList = headOfConvertList.left
	}
	return headOfConvertList

}

func ConvertNode(node *binaryTreeNode, lastNodeInList **binaryTreeNode) {
	if node == nil {
		return
	}
	current := node
	if current.left != nil {
		ConvertNode(current.left, lastNodeInList)
	}
	// 从递归环境中出来后，lastNodeInList 中存储的还是上层中的数据
	current.left = *lastNodeInList
	// 用来处理头节点
	if *lastNodeInList != nil {
		(*lastNodeInList).right = current
	}
	*lastNodeInList = current

	if current.right != nil {
		ConvertNode(current.right, lastNodeInList)
	}

}
