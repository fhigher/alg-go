package sorting

func InsertionSort(arr []int, n int) {
	var e, j int
	for i := 1; i < n; i++ {
		e = arr[i]
		for j = i; j > 0 && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

func InsertionSortRange(arr []int, l, r int) {
	var e, j int
	for i := l + 1; i <= r; i++ {
		e = arr[i]
		for j = i; j > l && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
