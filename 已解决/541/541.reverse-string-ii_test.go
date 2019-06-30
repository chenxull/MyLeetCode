package problem541

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	s      string
	k      int
	output string
}{
	{
		"abcdefg",
		2,
		"bacdfeg",
	},
}

func Test_reverseStr(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.output, reverseStr(tc.s, tc.k), "输入:%v", tc)

	}
}
