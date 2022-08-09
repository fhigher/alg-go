package array

import (
	"fmt"
	"sort"
	"testing"
)

var data [][]int = [][]int{
	{0, 1, 0, 3, 12},
	{1, 3, 12, 0, 0},
	{1, 3, 12, 0},
	{0, 0, 0, 3, 12},
	{0, 2, 2, 0, 2, 5, 0, 0, -1, 3, 12, 0},
}

func TestMoveZero(t *testing.T) {
	for _, d := range data {
		//fmt.Println("before: ", d)
		moveZero(d)
		fmt.Println("after: ", d)
	}
}

func TestRemoveElem(t *testing.T) {
	for _, d := range data {
		//fmt.Println("before: ", d)
		length := removeElem(d, 0)
		fmt.Println("after: ", d[:length])
	}
}

func TestRemoveDuplicates(t *testing.T) {
	for _, d := range data {
		sort.Ints(d)
		fmt.Println("before: ", d)
		length := removeDuplicates(d)
		fmt.Println("after: ", d[:length])
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	for _, d := range data {
		sort.Ints(d)
		fmt.Println("before: ", d)
		length := removeDuplicates2(d)
		fmt.Println("after: ", d[:length])
	}
}

func TestBinarySearch(t *testing.T) {
	for _, d := range data {
		sort.Ints(d)
		index := binarySearch(d, 5)
		fmt.Println("find 5 at :", index)
		index = binarySearchRecursive(d, 5)
		fmt.Println("find 5 at :", index)
	}
}

/* var data2 [][]int = [][]int{
	{2, 0, 2, 1, 1, 0},
	{2, 0, 1},
}

func TestSortColors(t *testing.T) {
	for _, d := range data2 {
		fmt.Println("before: ", d)
		sortColors(d)
		fmt.Println("after: ", d)
	}
} */
