package stack

type LinkStack struct {
	Data interface{}
	Next *LinkStack
}

// NewLinkStack 栈底
func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func (s *LinkStack) IsEmpty() bool {
	if s.Next == nil {
		return true
	}

	return false
}

// Push 头部插入
func (s *LinkStack) Push(data interface{}) {
	newData := new(LinkStack)
	newData.Data = data
	newData.Next = s.Next
	s.Next = newData
}

// Pop 头部取出，即先入先出
func (s *LinkStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	value := s.Next.Data
	s.Next = s.Next.Next

	return value
}

func (s *LinkStack) Length() int {
	next := s
	count := 0

	for next.Next != nil {
		count++
		next = next.Next
	}

	return count
}
