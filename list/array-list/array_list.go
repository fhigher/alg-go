package array_list

import (
	"errors"
	"fmt"
	"github.com/FHigher/algorithm/list"
)

type IArrayList interface {
	Size() int
	Cap() int
	IsEmpty() bool
	IsFull() bool
	Insert(index int, value interface{}) error // 在下标为index处插入value
	Append(value interface{})
	DeleteByValue(value interface{}) (int, error)
	DeleteByIndex(index int) (interface{}, error)
	Pop() (interface{}, error)
	Print()
	Clear()
	list.Iterable
}

type arrayList struct {
	elements []interface{}
	initSize int
}

var (
	ListIsEmpty     = errors.New("空列表")
	ListIsFull      = errors.New("列表已满")
	IndexOutOfRange = errors.New("索引超出范围")
)

// 初始化
func NewArrayList(size int) *arrayList {
	return &arrayList{
		elements: make([]interface{}, 0, size),
		initSize: size,
	}
}

// Clear
func (list *arrayList) Clear() {
	list.elements = make([]interface{}, 0, list.initSize)
}

// Append
func (list *arrayList) Append(value interface{}) {
	list.elements = append(list.elements, value)
}

// expand
func (list *arrayList) expand() {
	newList := make([]interface{}, 0, list.Size()*2)
	//copy(newList, list.elements) 此处copy没有发生作用，newList 仍然是空的，可能因为值类型是interface
	for _, e := range list.elements {
		newList = append(newList, e)
	}
	list.elements = newList
}

// Insert
func (list *arrayList) Insert(index int, value interface{}) error {

	if index < 0 || index >= list.Size() {
		return IndexOutOfRange
	}

	if list.IsFull() {
		// 扩容
		list.expand()
	}
	// 预分配一个空间，虽然右边是开空间，但是内存已经分配了，值为类型的默认值
	list.elements = list.elements[:list.Size()+1]

	for i := list.Size() - 2; i >= index; i-- {
		list.elements[i+1] = list.elements[i]
	}

	list.elements[index] = value

	return nil

}

// Size
func (list *arrayList) Size() int {
	return len(list.elements)
}

// Cap
func (list *arrayList) Cap() int {
	return cap(list.elements)
}

// Pop
func (list *arrayList) Pop() (value interface{}, err error) {

	size := list.Size()

	if size == 0 {
		return nil, ListIsEmpty
	}

	value = list.elements[size-1]
	list.elements = list.elements[:size-1]

	return
}

// IsEmpty
func (list *arrayList) IsEmpty() bool {

	if list.Size() == 0 {
		return true
	}

	return false
}

func (list *arrayList) IsFull() bool {
	if list.Size() == cap(list.elements) {
		return true
	}

	return false
}

// DeleteByValue
func (list *arrayList) DeleteByValue(value interface{}) (int, error) {

	delIndex := -1

	if list.IsEmpty() {
		return delIndex, ListIsEmpty
	}

	for i, e := range list.elements {
		//fmt.Println(i)，range遍历-到切片的实际长度则终止循环,即实际长度为实际分配的地址空间，
		// 而如果容量大于长度，则多余的容量还未分配空间，但是可以访问其下标而不能赋值
		if value == e {
			delIndex = i
			list.elements = append(list.elements[:delIndex], list.elements[delIndex+1:]...)
			break
		}
	}

	return delIndex, nil
}

// DeleteByValue
func (list *arrayList) DeleteByIndex(index int) (value interface{}, err error) {

	if list.IsEmpty() {
		return nil, ListIsEmpty
	}

	for i, e := range list.elements {
		if i == index {
			value = e
			list.elements = append(list.elements[:index], list.elements[index+1:]...)
			break
		}
	}

	return value, nil
}

// Print
func (list *arrayList) Print() {
	fmt.Printf("长度: %d, 容量：%d, 元素：%v\n", len(list.elements), cap(list.elements), list.elements)
}

func (list *arrayList) Iterator() list.Iterator {
	return NewArrayListIterator(list)
}
