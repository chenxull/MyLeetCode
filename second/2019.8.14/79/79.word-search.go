package problem79

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m == 0 {
		return false
	}

	if n == 0 {
		return false
	}

	// 因为所有位置的字母都可能作为单词的 首字母，所以要遍历数组中所有的元素
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i, j, index int) bool {
	if len(word) == index {
		return true
	}

	// 失败的条件
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	// 给即将判断的元素打上标记，避免重复计算
	temp := board[i][j]
	board[i][j] = 0

	if backtrack(board, word, i+1, j, index+1) ||
		backtrack(board, word, i-1, j, index+1) ||
		backtrack(board, word, i, j+1, index+1) ||
		backtrack(board, word, i, j-1, index+1) {
		return true
	}

	board[i][j] = temp
	return false
}
