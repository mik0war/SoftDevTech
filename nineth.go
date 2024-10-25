package main

import (
	"fmt"
	"math/rand/v2"
)

func minMax(array []int) (int, int) {

	//Помещаем максимальное значение int в переменную min
	min := int(^uint(0) >> 1)
	//Помещаем минимальное значение int в переменную max
	max := -min - 1

	for _, element := range array {

		if element > max {
			max = element
		}

		if element < min {
			min = element
		}
	}

	return min, max
}

func main() {

	for true {
		var n int
		fmt.Println("Введите размер массива")
		fmt.Scan(&n)

		array := make([]int, n)

		for i := 0; i < n; i++ {
			array[i] = rand.IntN(1000) - 500
		}
		fmt.Println("Source array: ", array)
		fmt.Println(minMax(array))
	}

}
