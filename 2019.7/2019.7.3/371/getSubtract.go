package problem371

func getSubtract(a int, b int) int {
	for a != 0 {
		a, b = ((^a)&b)<<1, a^b
	}
	return b
}
