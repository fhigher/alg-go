package array_list

import (
	"fmt"
	"github.com/FHigher/algorithm/utils"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	arrayList := NewArrayList(10)
	arrayList.Print()
}

func TestArrayList_Append(t *testing.T) {
	fmt.Print("最终数组: ")
	arrayList := arrayListAppend(50)
	arrayList.Print()
}

func arrayListAppend(size int) *arrayList {
	arrayList := NewArrayList(size)

	randNumArr := utils.GenerateRandNums(size)

	for _, n := range randNumArr {
		arrayList.Append(n)
	}
	return arrayList
}

var insertMap = map[int]string{
	0:  "bbb",
	5:  "ccc",
	10: "ggg",
	11: "eee",
}

func TestPop_Delete_Search(t *testing.T) {
	arrayList := arrayListAppend(10)
	//arrayList.Append("abc") // 此处使内部切片空间扩容为原来容量的2倍
	arrayList.Print()

	t.Run("pop", func(t *testing.T) {
		val, err := arrayList.Pop()
		if nil != err {
			fmt.Printf("pop %s\n", err.Error())
		}

		fmt.Println("pop value: ", val)
		arrayList.Print()
	})

	t.Run("delete_value", func(t *testing.T) {
		index, err := arrayList.DeleteByValue("ccc")
		if nil != err {
			fmt.Printf("delete %s\n", err.Error())
		}

		fmt.Println("delete value's index: ", index)
		arrayList.Print()
	})

	t.Run("delete_index", func(t *testing.T) {
		val, err := arrayList.DeleteByIndex(0)
		if nil != err {
			fmt.Printf("delete %s\n", err.Error())
		}

		fmt.Println("delete index's value: ", val)
		arrayList.Print()
	})

	t.Run("insert", func(t *testing.T) {
		for index, value := range insertMap {
			err := arrayList.Insert(index, value)
			if nil != err {
				fmt.Printf("insert bbb %s \n", err.Error())
				break
			}
			arrayList.Print()
		}
	})
}
