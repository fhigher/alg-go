package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Seqqueue struct {
	Maxsize int
	Data []int
	Head int
	Tail int
}

func NewSeqqueue(maxsize int) *Seqqueue {
	return &Seqqueue{
		Maxsize: maxsize,
		Data: make([]int, maxsize),
		Head: 0,
		Tail: 0,
	}
}

func (q *Seqqueue) IsFull() bool {
	return (q.Tail + 1) % q.Maxsize == q.Head
}

func (q *Seqqueue) IsEmpty() bool {
	return q.Tail == q.Head
}

func (q *Seqqueue) Size() int {
	return (q.Tail + q.Maxsize - q.Head) % q.Maxsize
}

func (q *Seqqueue) Push(e int) error{
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

func (q *Seqqueue) All() ([]int, error){
	size := q.Size()
	if 0 == size {
		return nil, errors.New("queue is empty")
	}

	return q.Data, nil
}

func (q *Seqqueue) GetFront() (front int, err error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	front = q.Data[(q.Head + 1) % q.Maxsize]
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

var Maxsize *int

func init() {
	Maxsize = flag.Int("maxsize", 5, "queue maxsize")
	flag.Parse()
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

func main() {

	queue := NewSeqqueue(*Maxsize)
	key := ""
	e := 0

	showMenu()

	for {
		fmt.Print("请输入选项：")
		fmt.Scanln(&key)
		switch key {
		case "menu":
			showMenu()
		case "push":
			fmt.Print("请输入入队元素：")
			fmt.Scanln(&e)
			if err := queue.Push(e); nil != err {
				fmt.Println(err.Error())
			} else {
				fmt.Println("入队OK")
			}
		case "pop":
			if out, err := queue.Pop(); nil != err {
				fmt.Println(err.Error())
			} else {
				fmt.Println("出队元素为：", out)
			}
		case "front":
			if front, err := queue.GetFront(); nil != err {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队头元素为：", front)
			}
		case "show":
			if all, err := queue.All(); nil != err {
				fmt.Println(err.Error())
			} else {
				fmt.Print("当前队列中元素为：")
				fmt.Println(all)
			}
		case "state":
			fmt.Println("当前队列状态：")
			queue.Status()
		case "clear":
			choice := ""
			for {
				fmt.Print("确定要清空队列? (y/n)：")
				fmt.Scanln(&choice)
				if "y" != choice && "n" != choice {
					fmt.Println("输入有误，请输入y/n")
				}

				if "y" == choice {
					queue.Clear()
					fmt.Println("队列已清空")
					break
				}

				if "n" == choice {
					break
				}
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("没有此菜单项")
		}
	}
}

