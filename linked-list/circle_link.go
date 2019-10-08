package main

import "fmt"

type CircleNode struct {
	No int
	Next *CircleNode
}

func InsertNode(head *CircleNode, node *CircleNode) {
	if head.Next == nil {
		head.No = node.No
		head.Next = head
		return
	}

	temp := head

	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
	node.Next = head
}

func ShowLink(head *CircleNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("%d --> ", temp.No)
		if temp.Next == head {
			break
		}

		temp = temp.Next
	}
}

func main() {
	head := &CircleNode{}


	InsertNode(head, &CircleNode{1, nil})
	InsertNode(head, &CircleNode{0, nil})
	InsertNode(head, &CircleNode{3, nil})
	InsertNode(head, &CircleNode{2, nil})
	ShowLink(head)
}
