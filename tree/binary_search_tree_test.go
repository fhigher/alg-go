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
	output(t, bst)
}

func TestBstIterator(t *testing.T) {
	node := NewNode(Key(15), Value("root"))
	bst := Bst{root: node}
	for _, d := range data {
		bst.InsertIterator(Key(d.k), Value(d.v))
	}
	output(t, bst)
}

func output(t *testing.T, bst Bst) {
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
	fmt.Println()

	fmt.Println("最小值： ", bst.Min().k)

	fmt.Println("最大值： ", bst.Max().k)

	fmt.Println("删除最小值前元素数量： ", bst.Size())
	bst.RemoveMin()

	fmt.Println("删除后最小值： ", bst.Min().k, "元素数量： ", bst.Size())
	fmt.Println()

	fmt.Println("删除最大值前元素数量： ", bst.Size())
	bst.RemoveMax()

	fmt.Println("删除后最大值： ", bst.Max().k, "元素数量： ", bst.Size())
	fmt.Println()

	fmt.Println("删除任意元素: ", 8)
	bst.Remove(Key(8))
	bst.MidOrderIterator()
	fmt.Println()
	fmt.Println("删除任意元素后的元素数量： ", bst.Size())
	fmt.Println()
}
