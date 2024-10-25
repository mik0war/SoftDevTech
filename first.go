package main

import (
	"fmt"
)

func addNumberDigits(number int32) int32 {
	var sum int32
	//Максимальное количество цифр в числе типа int32 - 10,
	//поэтому повторяем операцию 10 раз, чтобы не использовать цикл
	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	return sum
}

func main() {
	fmt.Println(addNumberDigits(2_147_483_647))
}
