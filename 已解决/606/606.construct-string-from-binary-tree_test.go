package problem606

import (
	//"fmt"
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	pre, in []int
	ans     string
}{

	{
		[]int{1, 2, 4, 3},
		[]int{4, 2, 1, 3},
		"1(2(4))(3)",
	},

	{
		[]int{1, 2, 4, 3},
		[]int{2, 4, 1, 3},
		"1(2()(4))(3)",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, tree2str(root), "输入:%v", tc)
		//fmt.Println("输出:", tree2str(root))
	}
}
