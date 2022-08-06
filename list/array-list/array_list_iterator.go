package array_list

import "github.com/FHigher/algorithm/list"

type ArrayListIterator struct {
	ArrayList  *arrayList
	FrontIndex int
	BackIndex  int
}

func NewArrayListIterator(arrayList *arrayList) list.Iterator {
	return &ArrayListIterator{
		ArrayList: arrayList,
		BackIndex: len(arrayList.elements) - 1,
	}
}

func (iter *ArrayListIterator) FrontHasNext() bool {
	if iter.FrontIndex == len(iter.ArrayList.elements) {
		return false
	}

	return true
}

func (iter *ArrayListIterator) FrontNext() interface{} {
	value := iter.ArrayList.elements[iter.FrontIndex]
	iter.FrontIndex++

	return value
}

func (iter *ArrayListIterator) BackHasNext() bool {
	if iter.BackIndex < 0 {
		return false
	}

	return true
}

func (iter *ArrayListIterator) BackNext() interface{} {

	value := iter.ArrayList.elements[iter.BackIndex]
	iter.BackIndex--

	return value
}
