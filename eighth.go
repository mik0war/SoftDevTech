package main

import (
	"errors"
	"fmt"
)

func reverceNumber(n int) int {
	if n < 0 {
		errors.New("Illegal argument")
	}

	result := 0

	for n > 0 {
		result = result*10 + (n % 10)
		n /= 10
	}

	return result
}

func isPalindrome(number int) bool {
	return number == reverceNumber(number)
}

func main() {

	for true {
		var n int
		fmt.Println("Введите число")
		fmt.Scan(&n)

		if isPalindrome(n) {
			fmt.Println("Палиндром")
		} else {
			fmt.Println("Не палиндром")
		}
	}

}
