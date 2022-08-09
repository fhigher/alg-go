package array

import "fmt"

type node struct {
	row, col, val int
}

const (
	rowCount = 5
	colCount = 4
)

func initArray(row, col int) (array [][]int) {
	array = make([][]int, row)
	for i := range array {
		array[i] = make([]int, col)
	}

	return
}

func printArray(a [][]int) {
	for _, row := range a {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

func transferSparse(a [][]int, row, col int) (sparseArray []node) {

	initNode := node{
		row: row,
		col: col,
		val: 0,
	}

	sparseArray = append(sparseArray, initNode)

	for i, row := range a {
		for j, value := range row {
			if value > 0 {
				tempNode := node{
					row: i,
					col: j,
					val: value,
				}
				sparseArray = append(sparseArray, tempNode)
			}
		}
	}

	return
}
