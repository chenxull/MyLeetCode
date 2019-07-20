package problem198

func robdp2(nums []int) int {
	var a, b int

	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(b+v, a)
		}
	}
	return max(a, b)
}
