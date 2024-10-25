package main

import (
	"fmt"
)

func digitsSqrt(number int) int {

	currentNum := 0
	for ok := true; ok; ok = (number/10 != 0) {
		for ; number > 0; number /= 10 {
			currentNum += number % 10
		}

		number = currentNum
		currentNum = 0
	}

	return number
}

func main() {

	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	fmt.Print(digitsSqrt(num))
}
