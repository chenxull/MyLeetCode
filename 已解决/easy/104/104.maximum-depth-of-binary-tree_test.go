package problem104

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	pre1, in1 []int
	dep       int
}{
	{
		pre1: []int{1, 3, 5, 2},
		in1:  []int{5, 3, 1, 2},
		dep:  3,
	},
}

func Test_Ok(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		fmt.Println("start")
		t1 := kit.PreIn2Tree(tc.pre1, tc.in1)
		ast.Equal(tc.dep, maxDepth(t1))
	}

}
