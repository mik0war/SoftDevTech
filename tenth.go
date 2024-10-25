package main

import (
	"fmt"
)

func isDivided5and3(num int) bool {
	return num%3 == 0 && num%5 == 0
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)

	if isDivided5and3(num) {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}
