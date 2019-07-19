package problem38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	intput int
	ans    string
}{
	{
		1,
		"1",
	},
	{
		4,
		"1211",
	},
	{
		5,
		"111221",
	},
}

func Test_ac38(t *testing.T) {
	ast := assert.New(t)

	for _, ts := range tsc {
		ast.Equal(ts.ans, countAndSay(ts.intput))
	}
}
