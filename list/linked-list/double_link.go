package main

import "fmt"

type Node struct {
	Name string
	Age  int
	Pre  *Node
	Next *Node
}

func NewNode(name string, age int) *Node {
	return &Node{
		Name: name,
		Age:  age,
	}
}

type DoubleLink struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewDoubleLink() *DoubleLink {
	return &DoubleLink{
		Head: NewNode("", 0),
	}
}

func (d *DoubleLink) DisorderInsert(node *Node) {
	if d.Head.Next == nil {
		d.Head.Next = node
		node.Pre = d.Head
		d.Tail = node
	} else {
		d.Tail.Next = node
		node.Pre = d.Tail
		d.Tail = d.Tail.Next
	}

	d.Length++
}

func (d *DoubleLink) OrderInsert(node *Node) {
	temp := d.Head
	for {
		if temp.Next == nil {
			temp.Next = node
			node.Pre = temp
			d.Tail = node
			d.Length++
			break
		} else if temp.Next.Age == node.Age && temp.Next.Name == node.Name {
			break
		} else if temp.Next.Age >= node.Age {
			temp.Next.Pre = node
			node.Next = temp.Next

			temp.Next = node
			node.Pre = temp
			d.Length++
			break
		}

		temp = temp.Next
	}
}

func (d *DoubleLink) DelNode(name string, age int) (node *Node) {
	temp := d.Head

	for {
		if temp.Next == nil {
			fmt.Println("没有找到要删除的元素")
			break
		}

		if temp.Next.Age == age && temp.Next.Name == name {
			node = temp.Next
			temp.Next = temp.Next.Next
			if temp.Next != nil {
				temp.Next.Pre = temp
			}
			d.Length--
			return
		}

		temp = temp.Next
	}

	return
}

// 顺序输出
func (d *DoubleLink) ShowList() {
	temp := d.Head

	fmt.Println("当前链表长度为：", d.Length)
	for {
		if temp.Next == nil {
			break
		}

		fmt.Printf("[%s(%d)] --> ", temp.Next.Name, temp.Next.Age)
		temp = temp.Next
	}
	fmt.Println()
	fmt.Println()
}

// 倒序输出
func (d *DoubleLink) BackShow() {
	temp := d.Tail

	for {
		if temp.Pre == nil {
			break
		}

		fmt.Printf("[%s(%d)]-->", temp.Name, temp.Age)
		temp = temp.Pre
	}
	fmt.Println()
	fmt.Println()
}
