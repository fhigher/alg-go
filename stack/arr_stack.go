package stack

import "errors"

type IStack interface {
	Size() int
	Cap() int
	Clear()
	Pop() (interface{}, error)
	Push(data interface{}) error
	IsFull() bool
	IsEmpty() bool
}

var (
	StackIsEmpty = errors.New("栈已空")
	StackIsFull  = errors.New("栈已满")
)

type Stack struct {
	elements []interface{}
	top      int
	initSize int
}

func NewStack(size int) IStack {
	return &Stack{
		elements: make([]interface{}, 0, size),
		top:      -1,
		initSize: size,
	}
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}

func (stack *Stack) Cap() int {
	return cap(stack.elements)
}

func (stack *Stack) Clear() {
	stack.elements = make([]interface{}, 0, stack.initSize)
	stack.top = -1
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}

	top := stack.elements[stack.top]
	stack.elements = stack.elements[:stack.top]
	stack.top--

	return top, nil
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return StackIsFull
	}

	stack.elements = append(stack.elements, data)
	stack.top++
	return nil
}

func (stack *Stack) IsFull() bool {
	return stack.Size() == stack.Cap()
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}
