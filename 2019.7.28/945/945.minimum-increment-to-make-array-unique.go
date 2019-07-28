package problem945

/*
 * @lc app=leetcode id=945 lang=golang
 *
 * [945] Minimum Increment to Make Array Unique
 */
//  本题所使用到的数学思想很精彩
func minIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}

	counts := [40001]int{}
	res := 0
	max := 0
	// 统计每个元素出现过多少次
	for _, v := range A {
		counts[v]++
		if max < v {
			max = v
		}
	}

	// 遍历到最大值处
	for i := 0; i < max; i++ {
		// 如果当前元素出现的次数小于等于1 ，说明其没有重复
		if counts[i] <= 1 {
			continue
		}
		// 记录下，将当前元素放在合适位置后，剩下还有多少个数字需要处理
		redundance := counts[i] - 1
		res += redundance
		// 将多余的元素向后移一位
		counts[i+1] += redundance
		// 这样处理之后，最后一位之前的所有元素都是独一无二的的。只有最后一位元素需要处理
	}

	// 处理完最后一位后，还剩下的多余数字。这些数字需要向后铺开
	redundance := counts[max] - 1
	res += (redundance + 1) * redundance / 2
	return res
}
