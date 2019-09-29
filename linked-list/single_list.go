package main

import (
	"fmt"
	"strings"
)

type PersonNode struct {
	Name string
	Age int
	Next *PersonNode
}

func NewPersonNode(name string, age int) *PersonNode {
	return &PersonNode{
		Name: name,
		Age: age,
	}
}

type SingleList struct {
	Head *PersonNode
	Tail *PersonNode
	Length int
}

// 无序尾部插入
func (list *SingleList) DisorderTailInsert(node *PersonNode) {
	if list.Head.Next == nil {
		list.Head.Next = node
		list.Tail = node
	} else {
		list.Tail.Next = node 	// 之后的结点往尾结点上加

		list.Tail = list.Tail.Next 	// 更新尾结点
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
		} else if temp.Next.Age >= node.Age && temp.Next.Name != node.Name{
			node.Next = temp.Next
			temp.Next = node
			list.Length++
			break
		} else if temp.Next.Age == node.Age && temp.Next.Name == node.Name{
			fmt.Println(node, "结点已存在")
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
		if 0 == strings.Compare(temp.Next.Name, name) {
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

		if 0 != strings.Compare(temp.Next.Name, name) {	// temp始终是正在比较结点的前一个结点
			temp = temp.Next	// temp后移
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
		Head:&PersonNode{},
	}
}

func main() {

	var singleList *SingleList

	singleList = NewSingleList()

	//fmt.Println("添加结点，无序：")
	//singleList.DisorderTailInsert(NewPersonNode("张三", 22))
	//singleList.DisorderTailInsert(NewPersonNode("李四", 19))
	//singleList.DisorderTailInsert(NewPersonNode("王天", 33))
	//singleList.DisorderTailInsert(NewPersonNode("玉玉", 10))
	//singleList.DisorderTailInsert(NewPersonNode("蛋蛋", 20))

	fmt.Println("添加结点，有序：")
	singleList.OrderInsert(NewPersonNode("张三", 22))
	singleList.OrderInsert(NewPersonNode("李四", 19))
	singleList.OrderInsert(NewPersonNode("王天", 33))
	singleList.OrderInsert(NewPersonNode("玉玉", 45))
	singleList.OrderInsert(NewPersonNode("王天", 33))
	singleList.OrderInsert(NewPersonNode("张衡", 33))
	singleList.ShowList()

	fmt.Println("查找结点：")
	person := singleList.Find("李四")
	unknown := singleList.Find("张飒")
	fmt.Println(person, unknown)
	fmt.Println()

	fmt.Println("删除结点：")
	firstPerson := singleList.Delete("张三")
	fmt.Println("被删除的结点", firstPerson)
	singleList.ShowList()
}


