package problem13

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
		"III",
		3,
	},
	{
		"IV",
		4,
	}, {
		"IX",
		9,
	}, {
		"LVIII",
		58,
	},
}

func Test_rotatedDigits(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.input, tc.output)
		ast.Equal(tc.output, romanToInt(tc.input))
		fmt.Println("输出", romanToInt(tc.input))
	}
}
