package problem383

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	ransomNote, magazine string
	output               bool
}{
	{
		"a",
		"b",
		false,
	},
	{
		"aa",
		"ab",
		false,
	},
	{
		"aa",
		"aab",
		true,
	},
	{
		"zxcvbnm",
		"asdfghjkqwertyuizxcvbnm",
		true,
	},
}

func Test_canConstruct(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.output, canConstruct(tc.ransomNote, tc.magazine), "输入:%v", tc)

	}
}
