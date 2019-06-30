package problems521

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	a   string
	b   string
	ans int
}{
	{
		"cdc",
		"cdc",
		-1,
	},
	{
		"abc",
		"qwe",
		3,
	},
	{
		"qwertyuiop",
		"asdfghjkl",
		10,
	},
}

func Test_findLUSlength(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.a, tc.b)
		ast.Equal(tc.ans, findLUSlength(tc.a, tc.b))
		fmt.Println("输出", findLUSlength(tc.a, tc.b))
	}

}

func Benchmark_findLUSlength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLUSlength(tc.a, tc.b)
		}
	}
}
