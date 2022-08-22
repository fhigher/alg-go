package sorting

import (
	"fmt"
	"testing"

	"github.com/fhigher/algorithm/heap"
	"github.com/fhigher/algorithm/utils"
)

func TestSort(t *testing.T) {
	length := 500000
	selectArr := utils.GenerateRandNums(length)
	insertArr := utils.GenerateRandNums(length)
	mergeArr := utils.GenerateRandNums(length)

	utils.RunningTime("selecting sort", SelectionSort, selectArr, length)
	utils.RunningTime("insertion sort", InsertionSort, insertArr, length)
	utils.RunningTime("merge sort", MergeSort, mergeArr, length)
}

type elem int

func (e elem) Compare(j heap.Elem) bool {
	return e > j.(elem)
}

func TestHeapSort(t *testing.T) {
	arr := utils.GenerateRandNums(20)
	arrElem := make([]heap.Elem, len(arr))

	for i := 0; i < len(arr); i++ {
		arrElem[i] = elem(arr[i])
	}
	fmt.Println("sort before: ", arrElem)
	//HeapSort(arrElem)
	HeapSortArray(arrElem)
	fmt.Println("sort after: ", arrElem)
}
