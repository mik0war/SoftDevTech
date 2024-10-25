package main

import (
	"fmt"
)

func pascalTriangle(level int) {
	previousLine := []int{}
	currentLine := []int{}

	for i := 0; i < level; i++ {
		currentLine = append(currentLine, 1)

		for i := 1; i < len(currentLine)-1; i++ {
			currentLine[i] = previousLine[i-1] + previousLine[i]
		}

		fmt.Println(currentLine)
		previousLine = []int{}
		previousLine = append(previousLine, currentLine...)
	}

}

func main() {

	for true {
		var n int
		fmt.Println("Введите уровень треугольника Паскаля")
		fmt.Scan(&n)

		pascalTriangle(n)
	}

}
