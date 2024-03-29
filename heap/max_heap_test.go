package heap

import (
	"fmt"
	"testing"
)

// Person ..
type Person struct {
	Name  string
	Money float64
}

// Compare ..
func (s *Person) Compare(j Elem) bool {
	return s.Money > j.(*Person).Money
}

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

func TestInsertAndDeleteMax(t *testing.T) {
	maxHeap := NewMaxHeap()

	for _, e := range data {
		maxHeap.Insert(e)
	}

	deleteMax(maxHeap)
}

func TestFromArrayAndDeleteMax(t *testing.T) {
	maxHeap := NewMaxHeap()

	arr := make([]Elem, len(data))
	for i := 0; i < len(data); i++ {
		arr[i] = data[i]
	}
	
	maxHeap.FromArray(arr)

	deleteMax(maxHeap)
}

func deleteMax(maxHeap *MaxHeap) {
	maxHeap.Print()

	for maxHeap.Size() > 0 {
		e, err := maxHeap.DeleteMax()
		if nil != err {
			fmt.Println(err)
		}

		fmt.Println(e)
	}

	maxHeap.Print()
}
