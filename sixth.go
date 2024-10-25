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

func main() {

	var n int
	for true {
		fmt.Println("Введите число:")
		fmt.Scan(&n)

		fmt.Println("Перевёрнутое число: ", reverceNumber(n))
	}

}
