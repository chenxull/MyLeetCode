package problem

// 不会写
/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Add and Search Word - Data structure design
 */
//  WordDictionary is
type WordDictionary struct {
	val  byte
	sons [26]*WordDictionary
	end  int
}

// Constructor is
/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord is
/** Adds a word into the data structure. */
func (w *WordDictionary) AddWord(word string) {
	for _, b := range word {
		idx := b - 'a'
		if w.sons[idx] == nil {
			w.sons[idx] = &WordDictionary{val: byte(b)}
		}
		w = w.sons[idx]
	}
	w.end++
}

// Search is
/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (w *WordDictionary) Search(word string) bool {
	for i, b := range word {
		if b != '.' {
			idx := b - 'a'
			// 字母不存在，返回 false
			if w.sons[idx] == nil {
				return false
			}
			w = w.sons[idx]
		} else {
			// 对于'.‘ 搜索的处理
			for _, son := range w.sons {
				if son == nil {
					continue
				}

				w = son
				// 这里有点不懂，有时间在深入理解一下
				if i == len(word)-1 {
					if w.end > 0 {
						return true
					}
					continue
				}

				if w.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
	}

	if w.end > 0 {
		return true
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
