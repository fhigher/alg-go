package tree

// Trie ..
type Trie struct {
	root *TrieNode
}

// NewTrie ..
func NewTrie(node *TrieNode) Trie {
	return Trie{node}
}

// Insert ..
func (t Trie) Insert(name string) {
	cur := t.root
	for _, v := range name {
		index := v - 'a'
		if cur.Children[index] == nil {
			cur.Children[index] = new(TrieNode)
		}

		cur = cur.Children[index]
	}

	cur.IsWord = true
}

// Find ..
func (t Trie) Find(needle string) bool {
	cur := t.root
	for _, v := range needle {
		index := v - 'a'
		if cur.Children[index] == nil {
			return false
		}

		cur = cur.Children[index]
	}

	return cur.IsWord
}

// TrieNode ..
type TrieNode struct {
	IsWord   bool
	Children [26]*TrieNode
}
