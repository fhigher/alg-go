package main

import (
	"github.com/fhigher/algorithm/sort/sorting"
	"github.com/fhigher/algorithm/utils"
)

func main() {
	length := 500000
	selectArr := utils.GenerateRandNums(length)
	insertArr := utils.GenerateRandNums(length)
	mergeArr := utils.GenerateRandNums(length)

	utils.RunningTime("selecting sort", sorting.SelectionSort, selectArr, length)
	utils.RunningTime("insertion sort", sorting.InsertionSort, insertArr, length)
	utils.RunningTime("merge sort", sorting.MergeSort, mergeArr, length)
}
