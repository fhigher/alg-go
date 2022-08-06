package main

import (
	"fmt"
	"strings"
)

type PersonNode struct {
	Name string
	Age  int
	Next *PersonNode
}

func NewPersonNode(name string, age int) *PersonNode {
	return &PersonNode{
		Name: name,
		Age:  age,
	}
}

type SingleList struct {
	Head   *PersonNode
	Tail   *PersonNode
	Length int
}

// 无序尾部插入
func (list *SingleList) DisorderTailInsert(node *PersonNode) {
	if list.Head.Next == nil {
		list.Head.Next = node
		list.Tail = node
	} else {
		list.Tail.Next = node // 之后的结点往尾结点上加

		list.Tail = list.Tail.Next // 更新尾结点
	}

	list.Length++
}

// 有序插入, 增序
func (list *SingleList) OrderInsert(node *PersonNode) {

	temp := list.Head

	for {
		if temp.Next == nil {
			temp.Next = node
			list.Length++
			break
		} else if temp.Next.Age == node.Age && temp.Next.Name == node.Name {
			fmt.Println(node, "结点已存在")
			break
		} else if temp.Next.Age >= node.Age {
			node.Next = temp.Next
			temp.Next = node
			list.Length++
			break
		}

		temp = temp.Next
	}
}

func (list *SingleList) Find(name string) *PersonNode {

	temp := list.Head
	for {
		if temp.Next == nil {
			break
		}
		if strings.Compare(temp.Next.Name, name) == 0 {
			return temp.Next
		}
		temp = temp.Next

	}

	return nil
}

func (list *SingleList) Delete(name string) (popNode *PersonNode) {

	temp := list.Head
	for {
		if temp.Next == nil {
			break
		}

		if strings.Compare(temp.Next.Name, name) != 0 { // temp始终是正在比较结点的前一个结点
			temp = temp.Next // temp后移
		} else {
			// 找到要删除的结点，先保存
			popNode = temp.Next
			// 删除
			temp.Next = temp.Next.Next
			list.Length--
			return
		}
	}

	return nil
}

func (list *SingleList) ShowList() {

	fmt.Println("链表长度为：", list.Length)
	temp := list.Head
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

func NewSingleList() *SingleList {
	return &SingleList{
		Head: &PersonNode{},
	}
}
