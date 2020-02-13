package sorting

import "github.com/FHigher/algorithm/sort"

func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, low, high int) {

	/*if low >= high {
		return
	}*/
	if high-low <= 15 { // optimize 2
		sort.InsertionSortRange(arr, low, high)
		return
	}

	mid := (low + high) / 2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	if arr[mid] > arr[mid+1] { // optimize 1
		merge(arr, low, mid, high)
	}
}

func merge(arr []int, low, mid, high int) {
	aux := make([]int, high-low+1)
	copy(aux, arr[low:high+1])
	i, j := low, mid+1
	for k := low; k <= high; k++ {
		if i > mid {
			arr[k] = aux[j-low]
			j++
		} else if j > high {
			arr[k] = aux[i-low]
			i++
		} else if aux[i-low] >= aux[j-low] {
			arr[k] = aux[j-low]
			j++
		} else {
			arr[k] = aux[i-low]
			i++
		}
	}
}
