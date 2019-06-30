package problem925

/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 *
 * https://leetcode.com/problems/long-pressed-name/description/
 *
 * algorithms
 * Easy (44.42%)
 * Total Accepted:    14.3K
 * Total Submissions: 32.5K
 * Testcase Example:  '"alex"\n"aaleex"'
 *
 * Your friend is typing his name into a keyboard.  Sometimes, when typing a
 * character c, the key might get long pressed, and the character will be typed
 * 1 or more times.
 *
 * You examine the typed characters of the keyboard.  Return True if it is
 * possible that it was your friends name, with some characters (possibly none)
 * being long pressed.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: name = "alex", typed = "aaleex"
 * Output: true
 * Explanation: 'a' and 'e' in 'alex' were long pressed.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: name = "saeed", typed = "ssaaedd"
 * Output: false
 * Explanation: 'e' must have been pressed twice, but it wasn't in the typed
 * output.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: name = "leelee", typed = "lleeelee"
 * Output: true
 *
 *
 *
 * Example 4:
 *
 *
 * Input: name = "laiden", typed = "laiden"
 * Output: true
 * Explanation: It's not necessary to long press any character.
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * name.length <= 1000
 * typed.length <= 1000
 * The characters of name and typed are lowercase letters.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */
func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	nameSize := len(name)
	typedSize := len(typed)

	i, j := 0, 0
	for i < nameSize && j < typedSize {
		c := name[i]
		need, pressed := 0, 0
		for i < nameSize && name[i] == c {
			need++
			i++
		}
		for j < typedSize && typed[j] == c {
			pressed++
			j++
		}
		if pressed < need {
			return false
		}
	}
	return i == nameSize && j == typedSize
}
