package offer34

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type binaryTreeNode struct {
	value int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func FindPath(root *binaryTreeNode, expectedSum int) {
	if root == nil {
		return
	}
	stack := arraystack.New()
	currentSum := 0
	findPathCore(root, stack, expectedSum, currentSum)
}

func findPathCore(root *binaryTreeNode, stack *arraystack.Stack, expectedSum, currentSum int) {
	currentSum += root.value
	stack.Push(root.value)

	// 判断是否符合要求且为叶子节点
	isLeaf := root.left == nil && root.right == nil
	if isLeaf && currentSum == expectedSum {
		it := stack.Iterator()
		for it.End(); it.Prev(); {
			fmt.Print(it.Value(), " ")
		}
		fmt.Print("\n")
	}

	// 如果不是叶节点，继续遍历左子树
	if root.left != nil {
		findPathCore(root.left, stack, expectedSum, currentSum)
	}

	if root.right != nil {
		findPathCore(root.right, stack, expectedSum, currentSum)
	}

	stack.Pop()
}
