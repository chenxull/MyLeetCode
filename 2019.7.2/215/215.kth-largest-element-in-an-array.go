package problem215

import "container/heap"

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

func findKthLargest(nums []int, k int) int {
	temp := highHeap(nums)
	h := &temp
	heap.Init(h)

	if k == 1 {
		return (*h)[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]

}

// highHeap 实现了 heap 的接口
type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}

//  大顶堆
func (h highHeap) Less(i, j int) bool {
	// h[i] > h[j] 使得 h[0] == max(h...)
	return h[i] > h[j]
}
func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	//  长度减一
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
