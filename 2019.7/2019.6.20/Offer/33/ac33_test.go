package problem33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ac33(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(true, VerifySquenceOfBST([]int{5, 7, 6, 9, 11, 10, 8}, 0, 6),
		"是二叉搜索树的后序遍历序列")
	ast.Equal(false, VerifySquenceOfBST([]int{7, 4, 6, 5}, 0, 3),
		"不是二叉搜索树的后序遍历序列")
}
