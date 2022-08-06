package sorting

func SelectionSort(arr []int, n int) {
	var minIndex int
	for i := 0; i < n; i++ {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
