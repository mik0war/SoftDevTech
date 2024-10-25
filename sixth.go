package main

import (
	"fmt"
)

func isOdd(number int) bool {
	return number%2 == 0
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)

	if isOdd(num) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}
