package main

import (
	"container/list"
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	// 初始化
	h := Constructor(N)
	for {
		var key string
		var value uint64
		n, _ := fmt.Scan(&key, &value)
		if n == 0 {
			break
		} else {
			// fmt.Println(key, value)
			h.Put(key, value)
		}
	}

}

// LRUCache is
type LRUCache struct {
	cap int
	l   *list.List               // doubly linked list
	m   map[string]*list.Element // hash table for checking if list node exists
}

// Pair is the value of a list node.
type Pair struct {
	key   string
	value uint64
}

// Constructor is
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[string]*list.Element, capacity),
	}
}

// Put is
func (c *LRUCache) Put(key string, value uint64) {
	// 如果数据在 cache 中,更新
	if node, ok := c.m[key]; ok {

		if node.Value.(*list.Element).Value.(Pair).value < value {
			// 将这个元素放到表首
			c.l.MoveToFront(node)
			// 更新 数据
			node.Value.(*list.Element).Value = Pair{key: key, value: value}
		}

	} else {
		// 元素不存在与 cache 中，先检查表是否满了
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			fmt.Println(c.l.Back().Value.(*list.Element).Value.(Pair).key, c.l.Back().Value.(*list.Element).Value.(Pair).value)

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
