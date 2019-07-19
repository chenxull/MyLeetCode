package problem51

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	res := [][]string{}
	cols := make([]bool, n)
	//  \ 对角线
	d1 := make([]bool, 2*n)

	//  / 对角线
	d2 := make([]bool, 2*n)

	board := make([]string, n)

	dfs(0, cols, d1, d2, board, &res)
	return res
}

func dfs(r int, cols, d1, d2 []bool, board []string, res *[][]string) {

	if r == len(board) {
		tmp := make([]string, r)
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}

	n := len(board)

	for c := 0; c < len(board); c++ {
		//  处于相同对角线的点，这些值 是相同的
		id1 := r - c + n
		id2 := 2*n - r - c - 1
		if !cols[c] && !d1[id1] && !d2[id2] {
			// 构造每一行的数据
			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}
			b[c] = 'Q'
			board[r] = string(b)
			//  占位标记用
			cols[c], d1[id1], d2[id2] = true, true, true

			dfs(r+1, cols, d1, d2, board, res)
			//  解除标记
			cols[c], d1[id1], d2[id2] = false, false, false

		}

	}
}
