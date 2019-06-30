package main

import "fmt"

/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 *
 * https://leetcode.com/problems/robot-return-to-origin/description/
 *
 * algorithms
 * Easy (70.67%)
 * Total Accepted:    147.9K
 * Total Submissions: 208.3K
 * Testcase Example:  '"UD"'
 *
 * There is a robot starting at position (0, 0), the origin, on a 2D plane.
 * Given a sequence of its moves, judge if this robot ends up at (0, 0) after
 * it completes its moves.
 *
 * The move sequence is represented by a string, and the character moves[i]
 * represents its ith move. Valid moves are R (right), L (left), U (up), and D
 * (down). If the robot returns to the origin after it finishes all of its
 * moves, return true. Otherwise, return false.
 *
 * Note: The way that the robot is "facing" is irrelevant. "R" will always make
 * the robot move to the right once, "L" will always make it move left, etc.
 * Also, assume that the magnitude of the robot's movement is the same for each
 * move.
 *
 * Example 1:
 *
 *
 * Input: "UD"
 * Output: true
 * Explanation: The robot moves up once, and then down once. All moves have the
 * same magnitude, so it ended up at the origin where it started. Therefore, we
 * return true.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "LL"
 * Output: false
 * Explanation: The robot moves left twice. It ends up two "moves" to the left
 * of the origin. We return false because it is not at the origin at the end of
 * its moves.
 *
 *
 */
func judgeCircle(moves string) bool {
	position := []int{0, 0}
	runes := []rune(moves)
	for _, step := range runes {
		if step == 'R' {
			position[1]++
		} else if step == 'L' {
			position[1]--
		} else if step == 'U' {
			position[0]++
		} else if step == 'D' {
			position[0]--
		} else {
			continue
		}
	}

	if position[0] == 0 && position[1] == 0 {
		return true
	}
	return false
}

func main() {
	move := "RULD"
	if judgeCircle(move) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
