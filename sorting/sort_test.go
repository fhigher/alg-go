package sorting

import (
	"testing"

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
