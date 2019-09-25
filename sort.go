package main

import (
	"github.com/FHigher/algorithm/sorting"
	"github.com/FHigher/algorithm/utils"
)

func main() {
	length := 50000
	//selectArr := []int{2, 6, 8, 7, 4, 3, 5, 1, 10, 9}
	insertArr := utils.GenerateRandNums(length)
	mergeArr := utils.GenerateRandNums(length)

	//utils.RunningTime("selecting sort", sorting.SelectionSort, selectArr, length)
	utils.RunningTime("insertion sort", sorting.InsertionSort, insertArr, length)
	utils.RunningTime("merge sort", sorting.MergeSort, mergeArr, length)
}
