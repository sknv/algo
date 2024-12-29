package leetcode

// https://leetcode.com/problems/implement-trie-prefix-tree/description/

type Trie struct {
	root *ByteTrieNode
}

func TrieConstructor() Trie {
	return Trie{
		root: NewByteTrieNode(26),
	}
}

func (this *Trie) Insert(word string) {
	curNode := this.root

	for i := range word {
		curSymbol := word[i]
		if curNode.children[curSymbol] == nil {
			curNode.children[curSymbol] = NewByteTrieNode(26)
		}

		curNode = curNode.children[curSymbol]
	}

	curNode.isWordEnd = true
}

func (this *Trie) Search(word string) bool {
	targetNode := this.searchNode(word)
	if targetNode == nil {
		return false
	}

	return targetNode.isWordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	targetNode := this.searchNode(prefix)

	return targetNode != nil
}

func (this *Trie) searchNode(word string) *ByteTrieNode {
	curNode := this.root

	for i := range word {
		curSymbol := word[i]
		if curNode.children[curSymbol] == nil {
			return nil
		}

		curNode = curNode.children[curSymbol]
	}

	return curNode
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
