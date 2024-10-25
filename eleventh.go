package main

import (
	"fmt"
)

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func main() {
	number := 5
	fmt.Println(factorial(number))
}
