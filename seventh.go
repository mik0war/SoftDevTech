package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	var year int
	fmt.Println("Введите число")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}
