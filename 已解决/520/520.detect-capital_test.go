package problem520

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output bool
}{
	{
		"USA",
		true,
	}, {
		"agC",
		false,
	}, {
		"FlaG",
		false,
	},
}

func Test_detectCapitalUse(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.input, tc.output)
		ast.Equal(tc.output, detectCapitalUse(tc.input))
		fmt.Println("输出", detectCapitalUse(tc.input))
	}
}
