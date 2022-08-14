package tree

import "fmt"

type Key int
type Value string

type Node struct {
	k Key
	v Value
	left *Node
	right *Node
}

func NewNode(k Key, v Value) *Node {
	return &Node{
		k,
		v,
		nil,
		nil,
	}
}

type Bst struct {
	root *Node
	size int
}

func (t *Bst) Size() int {
	return t.size
}

func (t *Bst) Root() *Node {
	return t.root
}

func (t *Bst) InsertRecursive(k Key, v Value) {
	t.root = t.recursiveInsert(t.root, k, v)
}

func (t *Bst) recursiveInsert(node *Node, k Key, v Value) *Node {
	if node == nil {
		t.size++
		return NewNode(k, v)
	}

	if node.k == k {
		node.v = v
	} else if node.k < k {
		node.right = t.recursiveInsert(node.right, k, v)
	} else {
		node.left = t.recursiveInsert(node.left, k, v)
	}

	return node
}

func (t *Bst) InsertIterator(k Key, v Value) {
	node := t.root
	pre := t.root
	for node != nil {
		if node.k == k {
			node.v = v
			break
		} else if node.k < k {
			pre = node
			node = node.right
		} else {
			pre = node
			node = node.left
		}
	}

	if node == nil {
		if pre.k > k {
			pre.left = NewNode(k, v)
		} 

		if pre.k < k {
			pre.right = NewNode(k, v)
		}

		t.size++
	}
}

func (t *Bst) Contain(k Key) bool {
	return t.contain(t.root, k)
}

func (t *Bst) contain(node *Node, k Key) bool {
	if node == nil {
		return false
	}

	if node.k == k {
		return true
	} else if node.k < k {
		return t.contain(node.right, k)
	} else {
		return t.contain(node.left, k)
	}
}

func (t *Bst) Search(k Key) *Value {
	return t.search(t.root, k)
}

func (t *Bst) search(node *Node, k Key) *Value {
	if node == nil {
		return nil
	}

	if node.k == k {
		return &node.v
	} else if node.k < k {
		return t.search(node.right, k)
	} else {
		return t.search(node.left, k)
	}
}

func (t *Bst) PreOrderIterator() {
	t.preOrder(t.root)
}

func (t *Bst) preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%v = %v, ", node.k, node.v)
		t.preOrder(node.left)
		t.preOrder(node.right)
	}
}

func (t *Bst) MidOrderIterator() {
	t.midOrder(t.root)
}

func (t *Bst) midOrder(node *Node) {
	if node != nil {
		t.midOrder(node.left)
		fmt.Printf("%v = %v, ", node.k, node.v)
		t.midOrder(node.right)
	}
}

func (t *Bst) LastOrderIterator() {
	t.lastOrder(t.root)
}

func (t *Bst) lastOrder(node *Node) {
	if node != nil {
		t.lastOrder(node.left)
		t.lastOrder(node.right)
		fmt.Printf("%v = %v, ", node.k, node.v)
	}
}