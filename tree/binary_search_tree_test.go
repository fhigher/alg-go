package tree

import (
	"fmt"
	"testing"
)

type entry struct {
	k int
	v string
}

var data = []entry{
	{10, "a"},
	{20, "b"},
	{8, "c"},
	{6, "d"},
	{7, "e"},
	{3, "f"},
	{40, "g"},
}

func TestBstRecursive(t *testing.T) {
	node := NewNode(Key(15), Value("root"))
	bst := Bst{root: node}
	for _, d := range data {
		bst.InsertRecursive(Key(d.k), Value(d.v))
	}

	t.Log("size: ", bst.Size())

	has := bst.Contain(Key(3))
	t.Log(has)

	v := bst.Search(Key(3))
	t.Log(*v)

	if !bst.Contain(Key(30)) {
		bst.InsertRecursive(Key(30), "new value")
	}

	t.Log("size: ", bst.Size())
	v = bst.Search(Key(30))
	t.Log(*v)

	fmt.Println("前序遍历： ")
	bst.PreOrderIterator()
	fmt.Println()

	fmt.Println("中序遍历： ")
	bst.MidOrderIterator()
	fmt.Println()

	fmt.Println("后序遍历： ")
	bst.LastOrderIterator()
	fmt.Println()

	fmt.Println("层序遍历： ")
	bst.TransverseOrder()
}

func TestBstIterator(t *testing.T) {
	node := NewNode(Key(15), Value("root"))
	bst := Bst{root: node}
	for _, d := range data {
		bst.InsertIterator(Key(d.k), Value(d.v))
	}

	t.Log("size: ", bst.Size())

	has := bst.Contain(Key(3))
	t.Log(has)

	v := bst.Search(Key(3))
	t.Log(*v)

	if !bst.Contain(Key(30)) {
		bst.InsertIterator(Key(30), "new value")
	}

	t.Log("size: ", bst.Size())
	v = bst.Search(Key(30))
	t.Log(*v)
	*v = "changed value"
	changed := bst.Search(Key(30))
	t.Log(*changed)

	fmt.Println("前序遍历： ")
	bst.PreOrderIterator()
	fmt.Println()
	fmt.Println("中序遍历： ")
	bst.MidOrderIterator()
	fmt.Println()
	fmt.Println("后序遍历： ")
	bst.LastOrderIterator()
	fmt.Println()
	fmt.Println("层序遍历： ")
	bst.TransverseOrder()
}
