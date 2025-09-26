package main

import (
	"fmt"
	"math"
)

func printMatrix(mat [][]int) {
	maxIntLen := len(fmt.Sprint(mat[0][0]))

	for _, col := range mat {
		for _, value := range col {
			fmt.Printf("%*d ", maxIntLen+1, value)
		}
		fmt.Print("\n")
	}
}

/*
	Дано послідовність чисел b1, ..., bn. Отримати квадратну матрицю порядку n, елементами якої є числа b1, ..., bn^n, розташовані за схемою:

9 8 4
7 5 3
6 2 1
*/
func main() {
	var err error
	var matrixSize int

	fmt.Print("Enter matrix size: ")
	_, err = fmt.Scan(&matrixSize)
	if err != nil {
		panic(err)
	}

	mat := make([][]int, matrixSize)
	for i := range mat {
		mat[i] = make([]int, matrixSize)
	}

	num := 1
	row := 0
	col := 0

	for i := 0; i < matrixSize*matrixSize; i++ {
		mat[matrixSize-row-1][matrixSize-col-1] = int(math.Pow(float64(num), float64(num)))
		num++

		if (row+col)%2 == 0 {
			if col == matrixSize-1 {
				row++
			} else if row == 0 {
				col++
			} else {
				row--
				col++
			}
		} else {
			if row == matrixSize-1 {
				col++
			} else if col == 0 {
				row++
			} else {
				col--
				row++
			}
		}
	}

	printMatrix(mat)
}
