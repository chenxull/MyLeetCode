package problem961

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input []int
	ouput int
}{
	{
		[]int{1, 2, 3, 3},
		3,
	},
	{
		[]int{2, 1, 2, 5, 3, 2},
		2,
	},
	{
		[]int{5, 1, 5, 2, 5, 3, 5, 4},
		5,
	},
}

func Test_repeatedNTimes(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ouput, repeatedNTimes(tc.input), "输入:%v", tc)
		fmt.Println("输入:", tc.input, "输出：", repeatedNTimes(tc.input))
	}
}
