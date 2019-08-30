package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		num[i] = tmp
	}
	fmt.Println(num)
	res := [][]int{}

	for i := range num {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && num[i] == num[i-1] {
			continue
		}

		l, r := i+1, len(num)-1

		for l < r {
			s := num[i] + num[l] + num[r]
			switch {
			case s < k:
				// 较小的 l 需要变大
				l++
			case s > k:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{num[i], num[l], num[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(num, l, r)
			}
		}
	}

	
	size := len(res)
	if size == 0 {
		fmt.Print(-1, -1, -1)
		return
	}
	for i := 0; i < size; i++ {
		fmt.Println(res[i][0], res[i][1], res[i][2])
	}

}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
