package problem36

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValsudorow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValsudocol(board, col) {
			return false
		}
	}

	for grid := 0; grid < 9; grid++ {
		if !isValsudoGrid(board, grid) {
			return false
		}
	}

	return true
}

func isValsudorow(board [][]byte, row int) bool {
	var num [10]bool
	for col := 0; col < 9; col++ {
		n := converByte2Num(board[row][col])

		if n < 0 {
			continue
		}
		if num[n] {
			return false
		}
		num[n] = true
	}
	return true
}

func isValsudocol(board [][]byte, col int) bool {
	var num [10]bool
	for row := 0; row < 9; row++ {
		n := converByte2Num(board[row][col])

		if n < 0 {
			continue
		}

		if num[n] {
			return false
		}
		num[n] = true
	}
	return true
}

func isValsudoGrid(board [][]byte, grid int) bool {
	var num [10]bool

	row := (grid / 3) * 3
	col := (grid % 3) * 3

	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			n := converByte2Num(board[row+drow][col+dcol])

			if n < 0 {
				continue
			}

			if num[n] {
				return false
			}

			num[n] = true
		}
	}

	return true
}

func converByte2Num(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
