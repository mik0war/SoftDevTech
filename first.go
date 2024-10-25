package main

import (
	"fmt"
)

func isPrime(number uint) uint {
	if number <= 1 {
		return 1
	}
	for i := uint(2); i < number; i++ {
		if number%i == 0 {
			return i
		}
	}

	return 0
}

func main() {
	var num uint
	fmt.Println("Введите число")
	fmt.Scan(&num)

	isPrime := isPrime(num)

	switch isPrime {
	case 0:
		fmt.Println("Число простое")
		break
	case 1:
		fmt.Println("Не простое (1 или 0)")
		break

	default:
		fmt.Printf("Число составное, наименьший делитель - %d", isPrime)
	}
}
