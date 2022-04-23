package leetc

type Trie struct {
	val   byte
	next  [26]*Trie
	isEnd bool
}

func Constructor208() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	bStr := []byte(word[:])
	parentTrie := this
	var curTire *Trie
	for _, v := range bStr {
		curTire = parentTrie.next[int(v-'a')]
		if curTire == nil {
			parentTrie.next[int(v-'a')] = new(Trie)
			curTire = parentTrie.next[int(v-'a')]
			curTire.val = v
			parentTrie = curTire
		} else {
			parentTrie = curTire
		}
	}
	curTire.isEnd = true
}

func (this *Trie) Search(word string) bool {
	bStr := []byte(word[:])
	parentTrie := this
	var curTire *Trie
	for _, v := range bStr {
		curTire = parentTrie.next[int(v-'a')]
		if curTire == nil {
			return false
		} else {
			parentTrie = curTire
		}
	}
	return curTire.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	bStr := []byte(prefix[:])
	parentTrie := this
	var curTire *Trie
	for _, v := range bStr {
		curTire = parentTrie.next[int(v-'a')]
		if curTire == nil {
			return false
		} else {
			parentTrie = curTire
		}
	}
	return true
}
