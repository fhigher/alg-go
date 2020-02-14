package tree

import (
	"testing"
)

var testData = []Transcation{
	[]byte("sdfkkkeke"),
	[]byte("232111"),
	[]byte(""),
	[]byte("jklj4"),
	[]byte("3f"),
	[]byte("44"),
	[]byte("343222"),
}

func TestNewMerkleTree(t *testing.T) {
	tree := NewMerkleTree(testData...)
	t.Log(tree.root)
}

func TestMerkleTree_PreOrderRange(t *testing.T) {
	tree := NewMerkleTree(testData...)
	tree.PreOrderRange(tree.root)
}
