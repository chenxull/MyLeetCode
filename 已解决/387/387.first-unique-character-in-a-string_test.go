package problem387

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output int
}{
	{
		"leetcode",
		0,
	},
	{
		"loveleetcode",
		2,
	},
}

func Test_fristUniqChar(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.output, firstUniqChar(tc.input), "输入:%v", tc)

	}
}
