package problem696

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output int
}{
	{
		"00110011",
		6,
	},
	{
		"10101",
		4,
	},
}

func Test_rotatedDigits(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.input, tc.output)
		ast.Equal(tc.output, countBinarySubstrings(tc.input))
		fmt.Println("输出", countBinarySubstrings(tc.input))
	}
}
