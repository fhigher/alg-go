package queue

import (
	"fmt"
	"os"
	"testing"
)

func TestCircleQueue(t *testing.T) {
	queue := NewSeqqueue(100)
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
