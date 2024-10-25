package main

import (
	"fmt"
	"math"
)

func getDigits(num int) []int {
	result := []int{}

	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}

	return result
}

func checkIsArmstrong(num int, digits []int) bool {
	pow := len(digits)
	digitsSum := 0
	for _, element := range digits {
		digitsSum += int(math.Pow(float64(element), float64(pow)))
	}

	return num == digitsSum
}

func main() {

	num := 0
	fmt.Scan(&num)

	if checkIsArmstrong(num, getDigits(num)) {
		fmt.Println("Число Армстронга")
	} else {
		fmt.Print("Не число Армстронга")
	}
}
