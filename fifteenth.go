package main

import (
	"fmt"
	"math/rand/v2"
)

func sumArray(array []int) int {
	var sum int
	for _, element := range array {
		sum += element
	}

	return sum
}

func main() {
	var arrSize int
	fmt.Println("Введите количество элементов в массиве: ")
	fmt.Scan(&arrSize)

	array := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		array[i] = rand.IntN(100)
	}
	fmt.Println(array)
	fmt.Println(sumArray(array))
}
