package problem647

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"abc",
		3,
	},

	{
		"aaa",
		6,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countSubstrings(tc.s), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countSubstrings(tc.s)
		}
	}
}
