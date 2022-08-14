package queue

import (
	"errors"
	"fmt"
)

type Seqqueue struct {
	Maxsize int
	Data    []int
	Head    int
	Tail    int
}

func NewSeqqueue(maxsize int) *Seqqueue {
	return &Seqqueue{
		Maxsize: maxsize,
		Data:    make([]int, maxsize),
		Head:    0,
		Tail:    0,
	}
}

func (q *Seqqueue) IsFull() bool {
	return (q.Tail+1)%q.Maxsize == q.Head
}

func (q *Seqqueue) IsEmpty() bool {
	return q.Tail == q.Head
}

func (q *Seqqueue) Size() int {
	return (q.Tail + q.Maxsize - q.Head) % q.Maxsize
}

func (q *Seqqueue) Push(e int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.Tail = (q.Tail + 1) % q.Maxsize
	q.Data[q.Tail] = e

	return nil
}

func (q *Seqqueue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	q.Head = (q.Head + 1) % q.Maxsize
	e := q.Data[q.Head]
	q.Data[q.Head] = 0
	return e, nil
}

func (q *Seqqueue) All() ([]int, error) {
	size := q.Size()
	if size == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.Data, nil
}

func (q *Seqqueue) GetFront() (front int, err error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	front = q.Data[(q.Head+1)%q.Maxsize]
	return
}

func (q *Seqqueue) Clear() {
	q.Head = 0
	q.Tail = 0
}

func (q *Seqqueue) Status() {
	fmt.Printf("    队列总容量: %d\n", q.Maxsize)
	fmt.Printf("    队列长度: %d\n", q.Size())
	fmt.Printf("    队列是否为空: %t\n", q.IsEmpty())
	fmt.Printf("    队列是否已满: %t\n", q.IsFull())
	fmt.Println()
}

func showMenu() {
	fmt.Println()
	fmt.Println("----- 菜	单")
	fmt.Println()
	fmt.Println("----- menu 	显示菜单")
	fmt.Println("----- push 	入队")
	fmt.Println("----- pop  	出队")
	fmt.Println("----- front 	获取队头")
	fmt.Println("----- show 	显示队列")
	fmt.Println("----- state	队列状态")
	fmt.Println("----- clear 	清空队列")
	fmt.Println("----- exit 	退出")
	fmt.Println()
}
