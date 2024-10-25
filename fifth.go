package main

import (
	"fmt"
)

var values = map[int]int{}

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	val, ok := values[n]
	if ok {
		return val
	}

	values[n] = fibonacci(n-1) + fibonacci(n-2)
	return values[n]
}

func main() {

	var n int
	for true {
		fmt.Println("Введите номер числа фибоначчи:")
		fmt.Scan(&n)

		fmt.Println("Число фиббоначи: ", fibonacci(n))
		fmt.Println("Сохранённые значения: ", values)
	}

}
