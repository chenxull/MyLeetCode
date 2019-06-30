package problem771

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	J     string
	S     string
	ouput int
}{
	{
		"aA",
		"aAAbbbb",
		3,
	},
	{
		"z",
		"ZZ",
		0,
	},
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ouput, numJewelsInStones(tc.J, tc.S), "输入:%v", tc)
		fmt.Println()
	}
}
