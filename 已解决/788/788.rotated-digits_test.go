package problem788

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  int
	output int
}{
	{
		10,
		4,
	},
	{
		9876,
		2257,
	},
	{
		10000,
		2320,
	},
}

func Test_rotatedDigits(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.input, tc.output)
		ast.Equal(tc.output, rotatedDigits(tc.input))
		fmt.Println("输出", rotatedDigits(tc.input))
	}
}
