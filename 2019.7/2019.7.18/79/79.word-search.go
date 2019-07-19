package problem79

/*

 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
func exist(board [][]byte, word string) bool {
	//  常规检查
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	//  i,j 表示 从二维数组字母的下标
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//  二维回溯
func backtracking(board [][]byte, word string, r, l, index int) bool {
	m, n := len(board), len(board[0])
	//终止条件
	if index == len(word) {
		return true
	}

	//  所有不符合条件的 都返回 false
	if r < 0 || l < 0 || m <= r || n <= l || word[index] != board[r][l] {
		return false
	}
	// 为了避免一个元素被重复访问，需要将访问的字母打上标记，在递归退出时回复。
	temp := board[r][l]
	board[r][l] = 0

	//  只要有一个返回了 true ，就可以了
	if backtracking(board, word, r+1, l, index+1) ||
		backtracking(board, word, r-1, l, index+1) ||
		backtracking(board, word, r, l+1, index+1) ||
		backtracking(board, word, r, l-1, index+1) {
		return true
	}

	board[r][l] = temp
	return false

}
