package offer43

//  参考连接 https://blog.csdn.net/yi_afly/article/details/52012593

func numberOf1Between1AndN(n int) int {
	if n < 1 {
		return 0
	}
	count := 0
	base := 1
	round := n

	for i := 1; i <= n; i++ {
		weight := round % 10
		round /= 10
		count += round * base
		if weight == 1 {
			count += n%base + 1
		} else if weight > 1 {
			count += base
		}
		base *= 10
	}
	return count
}
