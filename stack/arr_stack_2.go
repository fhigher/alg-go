package stack

import (
	"github.com/FHigher/algorithm/list"
	"github.com/FHigher/algorithm/list/array-list"
)

type IStack2 interface {
	Size() int
	Cap() int
	Clear()
	Pop() (interface{}, error)
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
	list.Iterable
}

type Stack2 struct {
	elements array_list.IArrayList
}

func NewStack2(arrayList array_list.IArrayList) IStack2 {
	return &Stack2{
		elements: arrayList,
	}
}

func (stack *Stack2) Size() int {
	return stack.elements.Size()
}

func (stack *Stack2) Cap() int {
	return stack.elements.Cap()
}

func (stack *Stack2) Clear() {
	stack.elements.Clear()
}

func (stack *Stack2) Pop() (interface{}, error) {
	return stack.elements.Pop()
}

func (stack *Stack2) Push(data interface{}) {
	stack.elements.Append(data)
}

func (stack *Stack2) IsFull() bool {
	return stack.elements.IsFull()
}

func (stack *Stack2) IsEmpty() bool {
	return stack.elements.IsEmpty()
}

func (stack *Stack2) Iterator() list.Iterator {
	return stack.elements.Iterator()
}
