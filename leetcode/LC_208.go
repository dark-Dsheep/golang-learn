package main

/*
*
208. 实现 Trie (前缀树) OVO
*/
type Trie struct {
	son   [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			cur.son[ch] = &Trie{}
		}
		cur = cur.son[ch]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			return false
		}
		cur = cur.son[ch]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, ch := range prefix {
		ch -= 'a'
		if cur.son[ch] == nil {
			return false
		}
		cur = cur.son[ch]
	}
	return true
}

func main() {

}
