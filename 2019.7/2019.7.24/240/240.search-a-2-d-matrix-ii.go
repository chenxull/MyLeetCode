package problem240

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	// 左下角开始遍历 ,右上角也可以使用这种方式
	// i, j := m-1, 0

	// for i >= 0 && j < n {
	// 	if matrix[i][j] == target {
	// 		return true
	// 	}

	// 	if matrix[i][j] > target {
	// 		i--
	// 	} else {
	// 		j++
	// 	}
	// }

	i, j := 0, n-1
	for j >= 0 && i < m {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
