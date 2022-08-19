package sorting

import (
	"github.com/fhigher/algorithm/heap"
)

func HeapSort(arr []heap.Elem) {
	length := len(arr)
	heap := heap.NewMaxHeap()
	for i := 0; i < length; i++ {
		heap.Insert(arr[i])
	}

	for i := length - 1; i >= 0; i-- {
		arr[i], _ = heap.DeleteMax()
	}
}
