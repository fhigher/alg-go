package heap

import (
	"fmt"
	"testing"
)

var data = []*Person{
	{
		Name:  "zhangsan",
		Money: 10000,
	},
	{
		Name:  "lisi",
		Money: 12000,
	},
	{
		Name:  "likui",
		Money: 12000,
	},
	{
		Name:  "hebei",
		Money: 13000,
	},
	{
		Name:  "caocao",
		Money: 9000,
	},
	{
		Name:  "gg",
		Money: 15000,
	},
	{
		Name:  "xiaoming",
		Money: 9500,
	},
	{
		Name:  "cc",
		Money: 11555,
	},
}

func TestInsertAndDelete(t *testing.T) {
	maxHeap := NewMaxHeap()

	for _, e := range data {
		maxHeap.Insert(e)
	}

	maxHeap.Print()

	for maxHeap.Size > 0 {
		e, err := maxHeap.Delete()
		if nil != err {
			fmt.Println(err)
		}

		fmt.Println(e)
	}

	maxHeap.Print()
}
