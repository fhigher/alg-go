package queue

import "errors"

type IQueue interface {
	Size() int
	Front() (interface{}, error)
	End() (interface{}, error)
	Enqueue(data interface{})
	Dequeue() (interface{}, error)
	Clear()
}

var (
	EmptyQueue = errors.New("队列为空")
)

type SimpleQueue struct {
	elements []interface{}
	size     int
}

func (q *SimpleQueue) Size() int {
	return q.size
}

func (q *SimpleQueue) Front() (interface{}, error) {
	if q.Size() == 0 {
		return nil, EmptyQueue
	}

	return q.elements[0], nil
}

func (q *SimpleQueue) End() (interface{}, error) {
	if q.Size() == 0 {
		return nil, EmptyQueue
	}

	return q.elements[q.Size()-1], nil
}

func (q *SimpleQueue) Enqueue(data interface{}) {
	q.elements = append(q.elements, data)
	q.size++
}

func (q *SimpleQueue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, EmptyQueue
	}

	data := q.elements[0]

	if q.Size() > 1 {
		q.elements = q.elements[1:q.Size()]
	} else {
		q.elements = make([]interface{}, 0)
	}

	q.size--

	return data, nil
}

func (q *SimpleQueue) Clear() {
	q.elements = make([]interface{}, 0)
	q.size = 0
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{
		elements: make([]interface{}, 0),
	}
}
