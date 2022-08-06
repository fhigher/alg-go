package main

import (
	"fmt"
	"testing"
)

func TestCircleNode(t *testing.T) {
	head := &CircleNode{}

	fmt.Println("添加节点")
	InsertNode(head, &CircleNode{1, nil})
	InsertNode(head, &CircleNode{0, nil})
	InsertNode(head, &CircleNode{3, nil})
	InsertNode(head, &CircleNode{2, nil})
	ShowLink(head)

	fmt.Println("删除节点")
	DelNode(head, 1)
	ShowLink(head)
	DelNode(head, 3)
	ShowLink(head)
	DelNode(head, 2)
	ShowLink(head)

	fmt.Println("添加节点")
	InsertNode(head, &CircleNode{5, nil})
	InsertNode(head, &CircleNode{7, nil})
	ShowLink(head)

	fmt.Println("删除节点")
	DelNode(head, 5)
	ShowLink(head)
}

func TestDoubleLink(t *testing.T) {
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

	fmt.Println("删除节点")
	d.DelNode("李四", 22)
	d.DelNode("张强", 44)
	d.ShowList()
}

func TestSingleList(t *testing.T) {

	singleList := NewSingleList()

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
