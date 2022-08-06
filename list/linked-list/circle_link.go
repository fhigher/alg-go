package main

import "fmt"

type CircleNode struct {
	No   int
	Next *CircleNode
}

func InsertNode(head *CircleNode, node *CircleNode) {
	if head.Next == nil {
		head.Next = node
		node.Next = head
		return
	}

	temp := head.Next

	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
	node.Next = head
}

func DelNode(head *CircleNode, no int) {
	temp := head
	if temp.Next == nil {
		fmt.Println("链表为空")
		return
	}

	for {
		if temp.Next.No == no {
			temp.Next = temp.Next.Next
			break
		}

		temp = temp.Next

		if temp.Next == head {
			break
		}
	}
}

func ShowLink(head *CircleNode) {
	temp := head.Next
	if temp == nil {
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

	fmt.Println()
}


