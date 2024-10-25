package main

import (
	"fmt"
	"math/rand/v2"
)

func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
			fmt.Println(array)
		}
	}

	return array
}

func main() {

	var n int
	fmt.Println("Введите размер массива")
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < len(array); i++ {
		array[i] = rand.IntN(100)
	}

	fmt.Println("Source array: ", array)
	fmt.Println("Sorted array: ", bubbleSort(array))

}
