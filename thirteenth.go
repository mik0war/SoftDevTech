package main

import (
	"bufio"
	"fmt"
	"os"
)

func createMatrix(matrixSize int) [][]int {

	matrix := make([][]int, matrixSize)

	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}

	return matrix
}

func getAliveNeighbours(x int, y int, matrix [][]int) int {
	result := 0
	for neighbourX := x - 1; neighbourX <= x+1; neighbourX++ {
		for neighbourY := y - 1; neighbourY <= y+1; neighbourY++ {
			if neighbourX < 0 || neighbourY < 0 || neighbourX >= len(matrix) ||
				neighbourY >= len(matrix[0]) || (neighbourX == x && neighbourY == y) {
				continue
			}
			if matrix[neighbourX][neighbourY] == 1 {
				result++
			}
		}
	}

	return result
}

func showMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Println(line)
	}
}

func makeStep(matrixSize int, matrix [][]int) [][]int {
	result := createMatrix(matrixSize)
	for x, line := range matrix {
		for y, cell := range line {
			result[x][y] = matrix[x][y]
			aliveNeighbours := getAliveNeighbours(x, y, matrix)
			if cell == 1 && !(aliveNeighbours == 3 || aliveNeighbours == 2) {
				result[x][y] = 0
				continue
			}

			if cell == 0 && aliveNeighbours == 3 {
				result[x][y] = 1
				continue
			}
		}
	}

	return result
}
func main() {

	const fieldSize = 10

	matrix := createMatrix(fieldSize)

	matrix[1][4] = 1
	matrix[2][2] = 1
	matrix[2][4] = 1
	matrix[3][3] = 1
	matrix[3][4] = 1

	for true {

		showMatrix(matrix)

		fmt.Println("Нажмите любую кнопку для продолжения")
		bufio.NewScanner(os.Stdin).Scan()

		matrix = makeStep(fieldSize, matrix)
	}
}
