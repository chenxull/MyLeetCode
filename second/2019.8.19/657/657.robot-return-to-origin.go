package problem657

import "strings"

/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */

func judgeCircle(moves string) bool {

	up := strings.Count(moves, "U")
	down := strings.Count(moves, "D")
	left := strings.Count(moves, "L")
	right := strings.Count(moves, "R")

	return up == down && left == right
}
