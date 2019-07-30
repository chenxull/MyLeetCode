package problem146

import "container/list"

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// LRUCache is
type LRUCache struct {
	cap int
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for checking if list node exists
}

// Pair is the value of a list node.
type Pair struct {
	key   int
	value int
}

// Constructor is
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get is
func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		c.l.MoveToFront(node)
		return val
	}
	return -1
}

// Put is
func (c *LRUCache) Put(key int, value int) {
	// 如果数据在 cache 中,更新
	if node, ok := c.m[key]; ok {
		// 将这个元素放到表首
		c.l.MoveToFront(node)
		// 更新 数据
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		// 元素不存在与 cache 中，先检查表是否满了
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			// 先删除 map 中的数据
			delete(c.m, idx)
			// 在删除双向链表中的数据
			c.l.Remove(c.l.Back())
		}
		// 将新的数据添加进来
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		// 放在链表的头部
		ptr := c.l.PushFront(node)
		// 数据放入 hashmap 中
		c.m[key] = ptr
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
