package problem925

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	name   string
	typed  string
	output bool
}{
	{
		"alex",
		"aaleex",
		true,
	},
	{
		"saeed",
		"ssaaedd",
		false,
	},
	{
		"leelee",
		"lleeelee",
		true,
	},
	{
		"laiden",
		"laiden",
		true,
	},
}

func Test_isLongPressedName(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.output, isLongPressedName(tc.name, tc.typed), "输入:%v", tc)

	}
}
