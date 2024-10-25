package main

import (
	"fmt"
	"math/rand/v2"
)

func reverce(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		temp := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = temp
	}

	return array
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
	fmt.Println(reverce(array))
}
