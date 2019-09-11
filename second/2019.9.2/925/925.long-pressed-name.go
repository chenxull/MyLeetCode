package problem925

/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 */
func isLongPressedName(name string, typed string) bool {
	// 统计相同元素出现的次数，通过比较次数来判断

	if name == typed {
		return true
	}

	namesize := len(name)
	typedsize := len(typed)

	i, j := 0, 0
	for i < namesize && j < typedsize {
		c := name[i]
		need, pressed := 0, 0
		for i < namesize && name[i] == c {
			need++
			i++
		}

		for j < typedsize && typed[j] == c {
			pressed++
			j++
		}
		if pressed < need {
			return false
		}
	}
	return i == namesize && j == typedsize
}
