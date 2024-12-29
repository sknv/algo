package leetcode

// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

type WordDictionary struct {
	root *ByteTrieNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{
		root: NewByteTrieNode(27),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curNode := this.root

	for i := range word {
		curSymbol := word[i]
		if curNode.children[curSymbol] == nil {
			curNode.children[curSymbol] = NewByteTrieNode(27)
		}

		curNode = curNode.children[curSymbol]
	}

	curNode.isWordEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchWordDictionaryNode(word, 0, this.root)
}

func searchWordDictionaryNode(word string, index int, node *ByteTrieNode) bool {
	if index == len(word) {
		return node.isWordEnd
	}

	char := word[index]
	if char == '.' {
		for _, child := range node.children {
			if searchWordDictionaryNode(word, index+1, child) {
				return true
			}
		}

		return false
	}

	if child, exists := node.children[char]; exists {
		return searchWordDictionaryNode(word, index+1, child)
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
