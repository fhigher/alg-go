package stack

import (
	"fmt"
	"testing"
)

func TestLinkStack(t *testing.T) {
	stack := NewLinkStack()

	// 入栈
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	fmt.Println(stack.Length())

	// 遍历
	node := stack.Next
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}

	// 出栈
	for data := stack.Pop(); data != nil; data = stack.Pop() {
		fmt.Println(data)
	}

	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Length())

}
