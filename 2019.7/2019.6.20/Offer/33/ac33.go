package problem33

func VerifySquenceOfBST(sequence []int, begin, end int) bool {
	if sequence == nil || begin > end {
		return false
	}

	root := sequence[end]
	i := begin
	// 左子树都小于根节点
	for ; i < end; i++ {
		if sequence[i] > root {
			break
		}
	}

	j := i
	// 右子树都大于根节点
	for ; j < end; j++ {
		if sequence[j] < root {
			return false
		}
	}

	// 判断左子树是否为二叉搜索后续遍历
	left := true
	// i - begin 是左子树的长度
	if 1 < i-begin {
		left = VerifySquenceOfBST(sequence, begin, i-1)
	}

	right := true

	// 判断右子树是否为二叉搜索后续遍历
	if 1 < end-i {
		right = VerifySquenceOfBST(sequence, i, end-1)
	}

	return right && left

}
