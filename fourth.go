package main

import (
	"fmt"
	"strings"
)

func mulMatrix(size int) [][]int {
	matrix := make([][]int, size)

	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}

	return matrix
}

func main() {

	var n int
	fmt.Println("Введите размер таблицы умножения")
	fmt.Scan(&n)
	matrix := mulMatrix(n)
	for i := 0; i < n; i++ {
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(matrix[i]), " ", "\t", -1), "[]"))
	}

}
