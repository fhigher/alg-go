package tree

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	names := []string{"jerry", "david", "lee"}
	needle := []string{"lee", "jerr", "joe"}

	trie := NewTrie(new(TrieNode))
	for _, name := range names {
		trie.Insert(name)
	}

	

	for _, name := range needle {
		if trie.Find(name) {
			fmt.Printf("find it: %s\n", name)
		} else {
			fmt.Printf("cannot find it: %s\n", name)
		}
	}

}
