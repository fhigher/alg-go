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

func HeapSortHeapify(arr []heap.Elem) {
	heap := heap.NewMaxHeap()

	heap.FromArray(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], _ = heap.DeleteMax()
	}
}

// 原数组空间进行堆排序， 此方式arr下标从0开始
func HeapSortArray(arr []heap.Elem) {
	for i := len(arr) / 2; i >= 0; i-- {
		shiftDown(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, i, 0)
	}
}

// 从index下标处开始shiftDown
func shiftDown(arr []heap.Elem, length int, index int) {

	// 2 × index + 2 是右孩子
	for 2*index+2 < length {
		max := 2*index + 1
		if 2*index+2 < length && arr[2*index+2].Compare(arr[max]) {
			max = 2*index + 2
		}

		if arr[max].Compare(arr[index]) {
			arr[index], arr[max] = arr[max], arr[index]
		} else {
			break
		}

		index = max
	}
}
