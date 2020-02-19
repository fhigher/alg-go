package array_list

import (
	"fmt"
	"github.com/FHigher/algorithm/list"
	"testing"
)

func CreateArrayListIterator(size int) list.Iterator {
	arrayList := ArrayListAppend(size)
	fmt.Println(arrayList.elements, len(arrayList.elements))
	iter := NewArrayListIterator(arrayList)
	return iter
}

func TestArrayListIterator_FrontNext(t *testing.T) {
	iter := CreateArrayListIterator(10)
	for iter.FrontHasNext() {
		value := iter.FrontNext()
		fmt.Printf("%v ", value)
	}

	fmt.Println()
}

func TestArrayListIterator_BackNext(t *testing.T) {
	iter := CreateArrayListIterator(10)
	for iter.BackHasNext() {
		value := iter.BackNext()
		fmt.Printf("%v ", value)
	}

	fmt.Println()
}
