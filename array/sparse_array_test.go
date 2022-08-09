package array

import (
	"fmt"
	"testing"
)

func TestSparseArray(t *testing.T) {
	var (
		array       [][]int
		sparseArray []node
		backArray   [][]int
	)
	array = initArray(rowCount, colCount)
	array[1][2] = 1
	array[2][3] = 2

	fmt.Println("原始数组：")
	printArray(array)

	fmt.Println("转化为稀疏数组：")
	sparseArray = transferSparse(array, rowCount, colCount)
	for _, row := range sparseArray {
		fmt.Printf("%d\t%d\t%d\n", row.row, row.col, row.val)
	}

	fmt.Println("复原为原始数组")
	backArray = initArray(rowCount, colCount)
	for _, row := range sparseArray[1:] {
		backArray[row.row][row.col] = row.val
	}
	printArray(backArray)
}
