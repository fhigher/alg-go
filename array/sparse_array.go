package array

import "fmt"

type node struct {
	row, col, val int
}

const (
	rowCount = 5
	colCount = 4
)

func main() {
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

	fmt.Println("稀疏数组：")
	sparseArray = transferSparse(array, rowCount, colCount)
	for _, row := range sparseArray {
		fmt.Printf("%d\t%d\t%d\n", row.row, row.col, row.val)
	}

	fmt.Println("原始数组")
	backArray = initArray(rowCount, colCount)
	for _, row := range sparseArray[1:] {
		backArray[row.row][row.col] = row.val
	}
	printArray(backArray)

}

func initArray(row, col int) (array [][]int) {
	array = make([][]int, row)
	for i, _ := range array {
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
