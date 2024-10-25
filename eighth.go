package main

import (
	"fmt"
)

func findMax(num1 int, num2 int, num3 int) int {
	isFirstBiggerSecond := num1 > num2
	isFirstBiggerThird := num1 > num3
	isSecondBiggerThird := num2 > num3

	if isFirstBiggerSecond && isFirstBiggerThird {
		return num1
	}

	if !isFirstBiggerSecond && isSecondBiggerThird {
		return num2
	}

	return num3
}

func main() {
	var num1, num2, num3 int
	fmt.Println("Введите числа")
	fmt.Scan(&num1, &num2, &num3)

	maxNum := findMax(num1, num2, num3)
	fmt.Println(maxNum)
}
