package problem380

import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */
type RandomizedSet struct {
	//  a 用来保存存入的元素
	a []int
	// map 的 k 就是 a 中元素，v 是 a 中元素的位置
	idx map[int]int
}

// Constructor is
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		a:   []int{},
		idx: make(map[int]int),
	}
}

// Insert is
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.idx[val]; ok {
		return false // 存在此元素
	}

	// 将新元素存储 a 中
	r.a = append(r.a, val)
	// 记录新加入元素在 a 中的位置
	r.idx[val] = len(r.a) - 1
	return true
}

// Remove is
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	// 先检测是否存在此元素
	if _, ok := r.idx[val]; !ok {
		return false // 不存在此元素
	}

	// 将 a 中最后的一个元素 移动到 待删除节点位置，并更新其 idx
	r.a[r.idx[val]] = r.a[len(r.a)-1]
	r.idx[r.a[len(r.a)-1]] = r.idx[val]

	// 删除最后一个元素
	r.a = r.a[:len(r.a)-1]
	delete(r.idx, val)
	return true
}

// GetRandom is
/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	return r.a[rand.Intn(len(r.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
