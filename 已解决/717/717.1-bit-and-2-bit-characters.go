package main

import "fmt"

/*
 * @lc app=leetcode id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 *
 * https://leetcode.com/problems/1-bit-and-2-bit-characters/description/
 *
 * algorithms
 * Easy (48.82%)
 * Total Accepted:    37.3K
 * Total Submissions: 76.1K
 * Testcase Example:  '[1,0,0]'
 *
 * We have two special characters. The first character can be represented by
 * one bit 0. The second character can be represented by two bits (10 or
 * 11).
 *
 * Now given a string represented by several bits. Return whether the last
 * character must be a one-bit character or not. The given string will always
 * end with a zero.
 *
 * Example 1:
 *
 * Input:
 * bits = [1, 0, 0]
 * Output: True
 * Explanation:
 * The only way to decode it is two-bit character and one-bit character. So the
 * last character is one-bit character.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * bits = [1, 1, 1, 0]
 * Output: False
 * Explanation:
 * The only way to decode it is two-bit character and two-bit character. So the
 * last character is NOT one-bit character.
 *
 *
 *
 * Note:
 * 1 .
 * bits[i] is always 0 or 1.
 *
 */
func isOneBitCharacter(bits []int) bool {
	if bits == nil || len(bits) == 0 {
		return false
	}

	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			continue
		}
		if bits[i] == 0 && i == len(bits)-1 {
			return true
		}
	}
	return false
}

func main() {
	array := []int{1, 1, 0}
	if isOneBitCharacter(array) {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}
