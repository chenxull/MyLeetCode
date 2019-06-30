package problem415

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num1 string
	num2 string
	ans  string
}{

	{
		"12",
		"3",
		"15",
	},

	{
		"2",
		"3",
		"5",
	},

	{
		"7",
		"8",
		"15",
	},

	// 可以有多个 testcase
}

func Test_addStrings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, addStrings(tc.num1, tc.num2), "输入:%v", tc)
		fmt.Println("结果为：", addStrings(tc.num1, tc.num2))
	}
}
