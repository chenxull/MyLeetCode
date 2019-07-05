package problem384

import "math/rand"

/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

type Solution struct {
	origNum, nums []int
}

func Constructor(nums []int) Solution {
	n := make([]int, len(nums))
	m := make([]int, len(nums))

	copy(n, nums)
	copy(m, nums)
	return Solution{origNum: n, nums: m}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origNum
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	i, j := len(this.nums), 0

	for i > 1 {
		j = rand.Intn(i)
		this.nums[i-1], this.nums[j] = this.nums[j], this.nums[i-1]
		i--
	}

	return this.nums

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
