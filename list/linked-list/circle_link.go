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

	return
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

func main() {
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
