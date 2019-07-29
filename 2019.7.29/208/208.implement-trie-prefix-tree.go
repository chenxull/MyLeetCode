package problem208

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

// Constructor is
func Constructor() Trie {
	return Trie{}
}

// Insert is
func (t *Trie) Insert(word string) {
	node := t
	size := len(word)

	for i := 0; i < size; i++ {
		idx := word[i] - 'a' // 获取字母的 ascii 值
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}
		node = node.sons[idx]
	}
	// 用来标记这是一个单词的结尾
	node.end++
}

// Search is
func (t *Trie) Search(word string) bool {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}
	if node.end > 0 {
		return true
	}
	return false

}

// StartsWith is
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
