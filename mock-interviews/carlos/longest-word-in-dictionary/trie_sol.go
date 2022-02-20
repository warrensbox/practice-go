type TrieNode struct {
	isWord   bool
	children map[rune]*TrieNode
}

func (this *TrieNode) hasAllPrefixes(word string) bool {
	it := this
	for _, ch := range word {
		if _, exists := it.children[ch]; !exists {
			return false
		}
		it = it.children[ch]
		if !it.isWord {
			return false
		}
	}
	return true
}

func (this *TrieNode) insert(word string) {
	it := this
	for _, ch := range word {
		if _, exists := it.children[ch]; !exists {
			it.children[ch] = &TrieNode{isWord: false, children: make(map[rune]*TrieNode)}
		}
		it = it.children[ch]
	}
	it.isWord = true
}

func longestWord(words []string) (best string) {
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	for _, word := range words {
		root.insert(word)
	}

	for _, word := range words {
		if len(word) < len(best) || (len(word) == len(best) &&
			word > best) {
			continue
		}
		if root.hasAllPrefixes(word) {
			best = word
		}
	}
	return
}