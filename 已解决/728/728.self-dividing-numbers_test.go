package problem728

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	left, right int
	ans         []int
}{
	{
		1,
		22,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
	},
}

func Test_selfDividingNumbers(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, selfDividingNumbers(tc.left, tc.right), "输入:%v", tc)

	}
}
