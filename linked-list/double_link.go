package main

import "fmt"

type Node struct {
	Name string
	Age int
	Pre *Node
	Next *Node
}

func NewNode(name string, age int) *Node {
	return &Node{
		Name: name,
		Age: age,
	}
}

type DoubleLink struct {
	Head *Node
	Tail *Node
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

func main() {
	d := NewDoubleLink()
	//fmt.Println("无序添加结点：")
	//d.DisorderInsert(NewNode("李四", 22))
	//d.DisorderInsert(NewNode("王伟", 11))
	//d.DisorderInsert(NewNode("张强", 44))

	fmt.Println("有序添加结点")
	d.OrderInsert(NewNode("李四", 22))
	d.OrderInsert(NewNode("王伟", 11))
	d.OrderInsert(NewNode("张强", 44))
	d.OrderInsert(NewNode("张强", 8))
	d.OrderInsert(NewNode("张强", 44))

	d.ShowList()
	d.BackShow()
}
